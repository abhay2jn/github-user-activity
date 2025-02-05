# GitHub User Activity

## Overview
The **GitHub User Activity** project is a command-line tool that allows you to fetch and display the recent activities of a specified GitHub user. It fetches the user's GitHub events via the GitHub API and prints a list of events along with the associated repositories. 

This tool provides an easy way to track what a GitHub user has been working on, whether it’s pushing code, opening issues, or other activities on GitHub.

## Features
- Fetch recent activities of a specified GitHub user.
- Displays activity type and associated repository for each event.
- Provides detailed information about the user's most recent actions on GitHub.

## Prerequisites
Before you can use the tool, ensure that you have the following installed:
- **Go** (version 1.15 or higher)
- **Cobra** package (for building command-line applications)

To install Cobra, run:
```bash
go get -u github.com/spf13/cobra@latest
```

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/abhay2jn/github-user-activity.git
   cd github-user-activity
   ```

2. Build the project:
   ```bash
   go build -o github-activity
   ```

3. Run the program:
   ```bash
   ./github-activity --user <username>
   ```

Replace `<username>` with the GitHub username you want to check. For example:
```bash
./github-activity --user abhay2jn
```

## Usage
### Command-line Usage
```bash
github-activity --user <username>
```
Where:
- `--user` or `-u`: The GitHub username to fetch activity for.

### Example
```bash
$ ./github-activity --user abhay2jn
abhay2jn recent Activity:
2025-02-05T08:23:34Z: PushEvent at github-user-activity
2025-02-05T08:18:20Z: PullRequestEvent at github-user-activity
```

## Explanation of Output
- The tool outputs the activity type (e.g., push, pull request) and the name of the repository where the activity occurred.
- The date and time of each activity are also displayed.

### Example Output:
```bash
abhay2jn recent Activity:
2025-02-05T08:23:34Z: PushEvent at github-user-activity
2025-02-05T08:18:20Z: PullRequestEvent at github-user-activity
```

- **PushEvent**: Indicates that a user pushed code to the repository.
- **PullRequestEvent**: Indicates that a user interacted with a pull request in the repository.

## Error Handling
- If there is an error fetching the data (e.g., invalid username or network issue), an error message will be shown.
- If the user has no recent activity, the tool will indicate that the user has no recent activity.

## Contributing
Feel free to fork this project and submit a pull request if you have improvements or bug fixes.

---

### Notes
- This tool utilizes the GitHub public API to fetch user events. Be mindful of rate limits when querying frequently.
- You can extend or modify the activity types as per the GitHub Events API documentation: [GitHub Events API](https://docs.github.com/en/rest/reference/activity#events).

---
