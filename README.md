# jira-api-resource

[![Build Status](https://travis-ci.org/TurnsCoffeeIntoScripts/jira-api-resource.svg?branch=master)](https://travis-ci.org/TurnsCoffeeIntoScripts/jira-api-resource)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/TurnsCoffeeIntoScripts/jira-api-resource/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/TurnsCoffeeIntoScripts/jira-api-resource/?branch=master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/e6ea2afc744d4fbf8bffc65e794155f4)](https://www.codacy.com/app/TurnsCoffeeIntoScripts/jira-api-resource?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=TurnsCoffeeIntoScripts/jira-api-resource&amp;utm_campaign=Badge_Grade)

***This ressource is still under construction. There is no stable release yet. Use at your own risk.***

This [Concourse](https://concourse-ci.org/) resource allows a pipeline to interface with a Jira REST API in order to manager (create/update/delete) tickets.

# Table of content
1. [Resource Type Configuration](#Resource-Type-Configuration)
2. [Source Configuration](#Source-Configuration)
    1. [Required Parameters Definition](#Required-Parameters-Definition)
    2. [Action Parameters Definition](#Action-Parameters-Definition)
        1. [Add Comment](#Add-Comment)

## Resource Type Configuration
``` yml
resource_types:
    - name: jira-api-resource
      type: docker-image
      source:
          repository: ... 
          tag: latest
```

## Source Configuration
``` yml
resources:
    - name: jira
      type: jira-resource
      source:
          url: https://...
          username: XXXX
          password: ((password-in-vault)
```

Firstly, here's a list of all required parameters:

### Required Parameters Definition
| Parameter  | Description                                                          |
|------------|----------------------------------------------------------------------|
| `url`      | The URL of the JIRA rest API to be used                              |
| `user`     | The username of the account used to connect with the Jira rest API   |
| `password` | The password of the specified user                                   |

### Action Parameters Definition
Now to specify the action to take specify one and only one of the actions listed bellow:

#### Add Comment
The configuration to use for add a comment to the specified ticket is `add-comment`.

**Example**:
``` yml
resources:
- name: qa
  type: git
  source:
    uri: ssh://git@git.abc.com:1234/test/dev/repo.git
    branch: develop
    private_key: ((git-key))
    tag_format: QA_{dateh}/v#
    tag_increment: num
    use_date: today
```