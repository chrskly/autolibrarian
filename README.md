# Autolibrarian

This tool will check the Puppetfile for each puppet environment on your
puppetmaster for use of a given combination of module and ref. If the module and
ref are used by an environment, the tool uses librarian-puppet to update the
module to the latest version.

### Configuration

The application searches for autolibrarian.conf in /etc/, then /usr/local/etc/.

Right now there is only one config setting ``puppet_environment_path``. This is
the directory on your puppetmaster which contains your puppet environments.

Example config
``puppet_environment_path: '/usr/local/etc/puppet/env'``

### Usage

``./autolibrarian -module <modulename> -ref <git-reference>``

Example

``./autolibrarian -module mycompany-myapp -ref develop``

### Triggering

You can use whatever makes sense for your environment to trigger the tool. E.g.,
git hook, Jenkins job, etc.
