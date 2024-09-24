# Github Foundations CLI
A command-line tool for the Github Foundations framework.

## Table of Contents

- [Usage](#usage)
    - [Generate](#generate)
    - [Import](#import)
    - [Check](#check)
    - [List](#list)
    - [Help](#help)
- [Installation](#installation)
    - [From releases](#from-releases)
        - [Linux](#linux)
        - [MacOS](#macos)
        - [Windows](#windows)
    - [From source](#from-source)

## Usage

There are a few main tools provided by the Github Foundations CLI:

```
Usage:
    github-foundations-cli [command]

Available Commands:
    gen         Generate HCL input for GitHub Foundations.
    import      Starts an interactive import process for resources in a Terraform plan.
    check       Perform checks against a Github configuration.
    list        List various resources managed by the tool.
    help        Help about any command.

Flags:
    -h, -- help     help for github-foundations-cli
```

### Generate

#### Generate using the Interactive mode

This command is used to generate HCL input for GitHub Foundations. This tool is used to generate HCL input for GitHub Foundations from state files output by terraformer.

```
Usage:
    github-foundations-cli gen <resource>
```

Where `<resource>` is one of the following:
- `repository_set`

Use `Shift + →` (right arrow) and `Shift + ←` (left arrow) to navigate through the questions.

Click on `Submit` to generate the HCL file.

### Import

This command will start an interactive process to import resources into Terraform state. It uses the results of a terraform plan to determine which resources are available for import.

```
Usage:
    github-foundations-cli import [module_path]

```

Where `<module_path>` is the path to the Terragrunt module to import.

### Check

Perform checks against a Github configuration and generate reports. This is used to validate the compliance stance of your GitHub configuration.

```
    Usage:
    github-foundations-cli check <org-slug>

```

Where `<org-slug>` is the organization slug to check.

### List

list various resources managed by the tool.


```
    Usage:
    github-foundations-cli list <resource> [options] [ProjectsDirectory|OrganzationsDirectory]

```

Where `<resource>` is one of the following:
- repos
- orgs


`[ProjectsDirectory]` is the path to the Terragrunt `Projects` directory when listing `repos`.

`[OrganzationsDirectory]` is the path to the Terragrunt `OrganzationsDirectory` directory when listing `orgs`.

`[options]` is a list of options to filter the list of resources. The options are:
- repos:
    - `--ghas`, `-g`    List repositories with GHAS enabled.

### Help

Display help for the tool.

## Installation

### From releases
Download the latest release from the [releases page](http:github.com/canada-ca/fondations-github-foundations/releases) and run the following commands:


#### Linux

**ADM64**
```
curl -LO https://github.com/canada-ca/fondations-github-foundations/releases/latest/download/github-foundations-cli_Linux_x86_64.tar.gz
tar -xzf github-foundations-cli_Linux_x86_64.tar.gz
chmod +x github-foundations-cli
sudo mv github-foundations-cli /usr/local/bin
```

**ARM64**
```
curl -LO https://github.com/canada-ca/fondations-github-foundations/releases/latest/download/github-foundations-cli_Linux_arm64.tar.gz
tar -xzf github-foundations-cli_Linux_arm64.tar.gz
chmod +x github-foundations-cli
sudo mv github-foundations-cli /usr/local/bin
```

#### MacOS

**ADM64**
```
curl -LO https://github.com/canada-ca/fondations-github-foundations/releases/latest/download/github-foundations-cli_Darwin_x86_64.tar.gz
tar -xzf github-foundations-cli_Darwin_x86_64.tar.gz
chmod +x github-foundations-cli
sudo mv github-foundations-cli /usr/local/bin
```

**ARM64**
```
curl -LO https://github.com/canada-ca/fondations-github-foundations/releases/latest/download/github-foundations-cli_Darwin_arm64.tar.gz
tar -xzf github-foundations-cli_Darwin_arm64.tar.gz
chmod +x github-foundations-cli
sudo mv github-foundations-cli /usr/local/bin
```

#### Windows

---
**i386**

1. Download the [latest release here](https://github.com/canada-ca/fondations-github-foundations/releases/download/v0.0.5/github-foundations-cli_Windows_i386.zip)

**ADM64**
1. Download the [latest release here](https://github.com/canada-ca/fondations-github-foundations/releases/download/v0.0.5/github-foundations-cli_Windows_i386.zip)

**ARM64**
1. Download the [latest release here](https://github.com/canada-ca/fondations-github-foundations/releases/download/v0.0.5/github-foundations-cli_Windows_i386.zip)
---

2. Unzip the package
3. Place the `github-foundations-cli.exe` executable in a directory of your choice, for example: `%USERPROFILE%\gh-foundations`

* **Add to Path (Optional):**
4. Right-click on "This PC" and select "Properties".
5. Click on "Advanced system settings".
6. Click on the "Environment Variables" button.
7. Under "System variables", find the "Path" variable and click "Edit".
8. Click "New" and add the following path: `%USERPROFILE%\gh-foundations` (replace with your chosen directory)
9. Click "OK" on all open windows to save the changes.

### From source
1.  Run `git clone git@github.com:canada-ca/fondations-github-foundations && cd github-foundations-cli/`
2.  Run `go mod download`
3.  Run `go build -v` for all providers OR build with one provider
