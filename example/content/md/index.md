# Pangolin CMS

Pangolin CMS is a content management system
for developers with the goal of simplicity,
configurability and minimalism.

## Simple to use

To create a website, you just need 3 folders and the
application itself:

- config
- content
- static

`config` contains the configuration files, while
you can put your website content into `content`.
Static files, like css and js, are placed into the
`static` folder.

The content is a collection of markdown or html files.
You can simply edit those files to edit your website.

After the setup, just start the cms and take a look
at your website!

## No database

The configuration is stored in YAML files and the
content of the website splittet into Markdown and
HTML files. You don't need a database.

## Extendability and configurability

When you need to modify the basic configuration
or want to add additional features, you can just
compile your own version.

the application is split into the code base (cms) and
the actual server (example). The example server
is ready to use. But you can also copy the example,
modify it like you wish and compile your own version.

You can still simply update the cms by updating the
version of the pangolin-cms version.
