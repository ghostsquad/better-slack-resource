# Slack Off

A [Concourse](https://concourse-ci.org/) resource for sending slack notifications


## Quick Start

Continue reading for `source` and `params` options.

```yaml
resource_types:
- name: slackoff
  type: docker-image
  source:
    repository: ghostsquad/slackoff
    tag: latest

resources:
- name: slack
  type: slackoff
  source:
    ...

jobs:
- name: example
  plan:
    - put: slack
      params:
        ...

```

## Source Configuration

- `url`: ***REQUIRED*** Incoming webhook url. See https://api.slack.com/incoming-webhooks
- `disable_put`: *optional* Convenience parameter for disabling all notifications during development/debugging. `default: false`
- `debug`: *optional* Prints the message to send as resource output. `default: false`
- `channel`: *optional* Overrides the default channel for the provided webhook url

## Put Params

- `template` - A [Go Template](https://golang.org/pkg/text/template/). of the slack message payload as described here in the [slack docs](https://api.slack.com/incoming-webhooks#advanced_message_formatting). Keep reading for more information on what is valid within the template.

- `template_file` - You may provide the template described above as a file instead of inline.

- `file_vars` - A JSON map, with the `key` being a string value you choose, the `value` is a path to a file.

    ```yaml
    file_vars:
      foo_key: path_to/file
      bar_key: path_to/another-file
    ```

    Files described are read and the content (with leading and trailing whitespace automatically trimmed) can be accessed within a template as follows:

    ```
    {{ .FileVars["foo_key"] }}
    ```

- `vars` - A JSON map of static keys & values that are available within a template.

    ```yaml
    vars:
      foo_key: foo value
      bar_key: bar value
    ```

    These values are accessed within a template as follows:

    ```
    {{ .Vars["foo_key"] }}
    ```

### Optional Params

- `channel`: *optional* Overrides the channel(s) set in the `source` configuration (if set), as well as the default channel for the webhook. #channel and @user forms are allowed. You can notify multiple channels separated by whitespace, like #channel @user.

- `channel_append`: *optional* Instead of overriding the `source` channel, this will add one more more channels to the list. #channel and @user forms are allowed. You can notify multiple channels separated by whitespace, like #channel @user. May be combined with `channel`.

- `channel_file`: *optional* File that contains a list of channels to send message to. May be used with `channel` and/or `channel_append`. The lists will be concatenated.

- `icon_url`: *optional* Override icon by providing URL to the image.

- `icon_emoji`: *optional* Override icon by providing emoji code (e.g. `:ghost:`)

### Dynamic Example

[***Compare to how this would be done with cfcommunity/slack-notification-resource***](#why-did-i-make-this)

```
...

resources:
- name: source
  type: git
  source:
    uri: git@github.com:example/example.git
    branch: example

- name: my-version
  type: semver
  source: { ... }

...

jobs:
- name: example
  plan:
  - get: source
  - get: my-version
    params: {bump: minor}
  - put: slack
    params:
      template: |
        {
            "attachments": [
                {
                  "fallback": "Build Started",
                  "color": "#439FE0",
                  "text": "Build Started",
                  "title": ":gear: Build started for {{ .Vars['PROJECT_NAME'] }}",
                  "title_link":  "{{ .Metadata.BuildUrl }}",
                  "fields": [
                    {
                      "title": "Project",
                      "value": "{{ .Vars['PROJECT_NAME'] }}",
                      "short": true
                    },
                    {
                      "title": "Revision",
                      "value": "{{ .FileVars['VERSION'] }}",
                      "short": true
                    }
                  ]
                }
            ]
        }
      file_vars:
        VERSION: my-version/version

      vars:
        PROJECT_NAME: my-project

```

## Get (Not Supported)

This is currently not supported, but I plan on seeing if it's possible to provide bot-like functionality within a concourse pipeline. Who knows!

## Check (Not Supported)

This is currently not supported, but I plan on seeing if it's possible to provide bot-like functionality within a concourse pipeline. Who knows!

## Why did I make this?

Well, I'll show up. Here's what my slack notifications used to look in my pipelines. I didn't want just static messages, I wanted rich, descriptive, dynamic messages. So I had to do this:

```
resource_types:
- name: slack-notification
  type: docker-image
  source:
    repository: cfcommunity/slack-notification-resource
    tag: latest

- name: metadata
  type: docker-image
  source:
    repository: olhtbr/metadata-resource

resources:
- name: slack-alert
  type: slack-notification
  source:
    url: ((slack-url))

- name: metadata
  type: metadata

jobs:
- name: example
  plan:
  - get: metadata
  - task: construct-starting-msg
      config:
        platform: linux
        image_resource:
          type: docker-image
          source:
            repository: ((generic-tools-image))
            tag: latest
        params:
          PROJECT_NAME: My Project
        inputs:
        - name: package-version-semver
        - name: metadata
        outputs:
        - name: starting-msg
        run:
          path: bash
          args:
            - -c
            - |
              set -euf -o pipefail
              export ATC_EXTERNAL_URL=$(cat metadata/atc_external_url)
              export BUILD_TEAM_NAME=$(cat metadata/build_team_name)
              export BUILD_PIPELINE_NAME=$(cat metadata/build_pipeline_name)
              export BUILD_JOB_NAME=$(cat metadata/build_job_name)
              export BUILD_ID=$(cat metadata/build_id)
              export BUILD_NAME=$(cat metadata/build_name)

              export PACKAGE_VERSION=$(cat ./package-version-semver/version)

              cat <<EOF > ./starting-msg/message.json
              [
                {
                  "fallback": "Build Started",
                  "color": "#439FE0",
                  "text": "Build Started",
                  "title": ":gear: Build started for ${PROJECT_NAME}",
                  "title_link":  "$ATC_EXTERNAL_URL/teams/$BUILD_TEAM_NAME/pipelines/$BUILD_PIPELINE_NAME/jobs/$BUILD_JOB_NAME/builds/$BUILD_NAME",
                  "fields": [
                    {
                      "title": "Project",
                      "value": "${PROJECT_NAME}",
                      "short": true
                    },
                    {
                      "title": "Revision",
                      "value": "${PACKAGE_VERSION}",
                      "short": true
                    }
                  ]
                }
              ]
              EOF

              cat ./starting-msg/message.json
  - put: slack-alert
    params:
      channel: "#concourse"
      username: concourse
      attachments_file: starting-msg/message.json
```

That's `59` lines of code for 1 slack notification, this is also divided between 3 steps (including metadata). Good Grief!!!

## Development

The docker container takes care of everything, see that for more instructions

```bash
docker build -t slack-off .
```

### Reference

https://api.slack.com/methods/chat.postMessage
