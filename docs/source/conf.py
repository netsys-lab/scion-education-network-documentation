# Configuration file for the Sphinx documentation builder.

# -- Project information

project = 'SCIERA'
copyright = '2024, OVGU / ETHZ'
author = 'OVGU Magdeburg, ETH Zurich'

release = '1.0'
version = '1.0.0'

# -- General configuration

extensions = [
    'sphinx.ext.duration',
    'sphinx.ext.doctest',
    'sphinx.ext.autodoc',
    'sphinx.ext.autosummary',
    'sphinx.ext.intersphinx',
    # Based on https://stackoverflow.com/a/73210637/817736 :
    'sphinx_rtd_size',
]

# Based on https://stackoverflow.com/a/73210637/817736 :
sphinx_rtd_size_width = "90%"

intersphinx_mapping = {
    'python': ('https://docs.python.org/3/', None),
    'sphinx': ('https://www.sphinx-doc.org/en/master/', None),
}
intersphinx_disabled_domains = ['std']

templates_path = ['_templates']

# -- Options for HTML output

html_theme = 'sphinx_rtd_theme'

# -- Options for EPUB output
epub_show_urls = 'footnote'
