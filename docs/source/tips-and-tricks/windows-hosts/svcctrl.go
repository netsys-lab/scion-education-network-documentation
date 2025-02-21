package main

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/eventlog"
	"golang.org/x/sys/windows/svc/mgr"
)

const ERROR_INVALID_NAME = 123

type Service struct {
	Name    string
	Exepath string
	Desc    string
	Config  string
	Logfile string
}

var services = []Service{
	{
		Name:    "SCION Daemon",
		Exepath: os.ExpandEnv("$ProgramFiles\\scion\\daemon.exe"),
		Desc:    "Facilitates communication of SCION applications with the control service.",
		Config:  os.ExpandEnv("$ProgramData\\scion\\config\\sd.toml"),
		Logfile: os.ExpandEnv("$ProgramData\\scion\\log\\sd.log"),
	},
}

func main() {
	var err error
	if len(os.Args) != 2 {
		fmt.Println("invalid number of arguments")
		return
	}
	if os.Args[1] == "install" {
		fmt.Println("Installing service")
		err = installServices()
	} else if os.Args[1] == "remove" {
		fmt.Println("Removing service")
		err = removeServices()
	} else if os.Args[1] == "start" {
		fmt.Println("Starting SCION")
		err = startServices()
	} else if os.Args[1] == "stop" {
		fmt.Println("Stopping SCION")
		err = stopServices()
	}
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func installServices() error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	for _, service := range services {
		c := mgr.Config{
			StartType:   mgr.StartManual,
			Description: service.Desc,
		}
		s, err := m.CreateService(service.Name, service.Exepath, c,
			"--config", service.Config,
			"--logfile", service.Logfile)
		if err != nil {
			if errno, ok := err.(syscall.Errno); ok {
				if uintptr(errno) == ERROR_INVALID_NAME {
					fmt.Printf("service %s already exists\n", service.Name)
					continue
				}
			}
			return err
		}
		defer s.Close()

		err = eventlog.InstallAsEventCreate(service.Name, eventlog.Error|eventlog.Warning|eventlog.Info)
		if err != nil {
			s.Delete()
			return fmt.Errorf("installing event log source failed: %v", err)
		}
	}

	return nil
}

func removeServices() error {
	return applyToServices(func(s *mgr.Service) error {
		err := eventlog.Remove(s.Name)
		if err != nil {
			return fmt.Errorf("removing event log source failed: %v", err)
		}
		err = s.Delete()
		if err != nil {
			return err
		}
		return nil
	})
}

func startServices() error {
	return applyToServices(func(s *mgr.Service) error {
		return s.Start()
	})
}

func stopServices() error {
	return applyToServices(func(s *mgr.Service) error {
		_, err := s.Control(svc.Stop)
		return err
	})
}

func applyToServices(f func(*mgr.Service) error) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	for _, service := range services {
		s, err := m.OpenService(service.Name)
		if err != nil {
			fmt.Printf("service %s not found\n", service.Name)
			continue
		}
		defer s.Close()
		err = f(s)
		if err != nil {
			return err
		}
	}
	return nil
}
