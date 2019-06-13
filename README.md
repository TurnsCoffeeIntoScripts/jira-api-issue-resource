# jira-api-resource  

| Version       | Build status | Scrutinizer |
|---------------|--------------|-------------|
| `0.0.1-rc.3`  | [![Build Status](https://travis-ci.org/TurnsCoffeeIntoScripts/jira-api-resource.svg?branch=master)](https://travis-ci.org/TurnsCoffeeIntoScripts/jira-api-resource) | [![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/TurnsCoffeeIntoScripts/jira-api-resource/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/TurnsCoffeeIntoScripts/jira-api-resource/?branch=master) |

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/e6ea2afc744d4fbf8bffc65e794155f4)](https://www.codacy.com/app/TurnsCoffeeIntoScripts/jira-api-resource?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=TurnsCoffeeIntoScripts/jira-api-resource&amp;utm_campaign=Badge_Grade)   
<sub>Project certification on default branch (master)</sub>

***This ressource is still under construction. There is no stable release yet. Use at your own risk.***

This [Concourse](https://concourse-ci.org/) resource allows a pipeline to interface with a Jira REST API in order to manage (create/update/delete) tickets.

# Table of content
1. [Resource Type Configuration](#Resource-Type-Configuration)
2. [Source Configuration](#Source-Configuration)
    1. [Required Parameters Definition](#Required-Parameters-Definition)
    2. [Action Parameters Definition](#Action-Parameters-Definition)
        1. [Add Comment](#Comment)
        2. [Add Label](#Add-Label)
    3. [Issue definition methods](#Issue-definition-methods)
        1. [Single Issue](#Single-issue)
        2. [List of issues](#List-of-issues)
        3. [Custom Script](#Custom-script)
3. [Behavior](#Behavior)

## Resource Type Configuration
``` yml
resource_types:
    - name: jira-api-resource
      type: docker-image
      source:
          repository: turnscoffeeintoscripts/jira-api-resource
          tag: latest
```

## Source Configuration
``` yml
resources:
    - name: jira
      type: jira-api-resource
      source:
          url: https://...
          username: XXXX
          password: ((password-in-vault)
          
          # Use only one of the next three parameters
          issue-id: ABC-123
          issue-list: ABC-123,ABC-234,ABC-345
          issue-script: /path/to/script/script.sh       
```

Firstly, here's a list of all required parameters:

### Required Parameters Definition

| Parameter      | Description                                                                       |
|----------------|-----------------------------------------------------------------------------------|
| `url`          | The URL of the JIRA rest API to be used                                           |
| `user`         | The username of the account used to connect with the Jira rest API                |
| `password`     | The password of the specified user                                                |

Next, use **one and only one** of the following parameters.

| Parameter      | Description                                                                       |
|----------------|-----------------------------------------------------------------------------------|
| `issue-id`     | The unique identifier of the Jira issue                                           |
| `issue-list`   | A list of all the Jira issue's unique identifier                                  |
| `issue-script` | Filename containing a script that must returns a single or multiple Jira issue(s) |

### Action Parameters Definition
Now to specify the action to take specify (set to 'true') one and only one of the actions listed bellow:

| Action (flags)     | Description               |
|--------------------|---------------------------|
| `comment`          | Add a comment on a ticket |
| `add-label`        | Add a label on a ticket   |

#### Comment
The configuration to use to add a comment to the specified ticket(s) is `comment`.

**Example**:
``` yml
resources:
    - name: jira
      type: jira-api-resource
      source:
          url: https://...
          username: XXXX
          password: ((password-in-vault)
          
          issue-list: ABC-123,ABC-234,ABC-345
          
          # The action to take on specified issue(s)
          comment: true
          body: This a comment made from a concourse resource
```

#### Add Label
The configuration to use to add a label to the specified ticket(s) is `add-label`.

**Example**:
```yml
resources:
    - name: jira
      type: jira-api-resource
      source:
          url: https://...
          username: XXXX
          password: ((password-in-vault)
          
          issue-list: ABC-123,ABC-234,ABC-345
          
          # The action to take on specified issue(s)
          add-label: true
          label: LABEL_XYZ
    
```

### Issue definition methods
With this resource, there are 3 possible ways of defining which tickets will be accessed or modified.
#### Single issue
#### List of issues
#### Custom Script

## Behavior