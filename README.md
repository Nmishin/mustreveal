# mustreveal
A tool to reveal (decrypt) obscured rclone configuration.

By default, in the rclone config file, human-readable passwords are obscured.
Obscuring them is done by encrypting them and writing them out in base64.

But rclone does not include command for reveal (decrypt) obscured configuration.
If you need to view obscured passwords, there is no simple option available.

The mustreveal tool allows you to easily view hidden fields in the configuration.

Currently, only the `secret` and `password` fields are accessible.

**Please note:** This tool is not intended to decrypt config files encrypted using `rclone config encryption` command.
