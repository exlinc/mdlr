# Using the `mdlr` tool

This guide discusses a few common ways to use `mdlr` for git dependency management

# Example workflow

1.  `cd project/directory/ # Enter the project directory`
2.  `mdlr init`
3.  `mdlr add --name depname --path deps/mydep --url https://github/org/mydep.git && echo '\ndeps/mydep/' >> .gitignore`
4.  `mdlr import -f # Reset the module forcefully (wipe changes, if any) and then import it at the version in the mdlr.yml file`
5.  `mdlr list # List the modules`
6.  `mdlr status # Get the status overview`
7.  `mdlr update -f # Reset the module forcefully (wipe changes, if any) and then update it and write the new update to the mdlr.yml file`
8.  `vim mdlr.yml # View/edit the mdlr.yml file`

# Create a new mdlr project

In the project directory, run `mdlr init`

# Import modules for a mdlr project

In the project directory, run `mdlr import -f`

# .gitignore

Make sure that **all** of the module paths are in the `.gitignore` file -- otherwise git will freak out!

# Commands overview

- `help`: get a help overview
- `init`: generate a mdlr.yml file in the directory
- `list`: list the current modules
- `add`: add a module to the mdlr.yml file
- `remove`: remove a module
- `import`: import a module
- `update`: update a module
- `status`: get the status for the mdlr.yml or invidual modules
