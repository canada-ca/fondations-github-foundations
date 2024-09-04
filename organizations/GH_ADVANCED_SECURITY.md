([Back](../README.md#included-tools))

# GitHub Advanced Security (GHAS) checks

The `GHAS Checks` tool is automatically installed on all GitHub Foundations organizations.

It runs [GitHub Advanced Security (GHAS)](https://docs.github.com/en/get-started/learning-about-github/about-github-advanced-security) scans for eligible repositories that have it enabled and reports all of the GHAS scans in one report. It is implemented as a GitHub Action that runs on a cron schedule (02:00 daily), and works by checking for any repositories that have GHAS enabled and running a scan on them.

## Changing the Schedule

The schedule for the `GHAS Checks` tool can be changed by modifying the `schedule` field in the `.github/workflows/ghas-policy-check.yml` file.

```yaml
on:
  schedule:
    - cron: '0 2 * * *'
```

The schedule is set to run at 02:00 daily by default.
To learn how cron schedules work, see the [GitHub Actions documentation](https://docs.github.com/en/actions/writing-workflows/choosing-when-your-workflow-runs/events-that-trigger-workflows#schedule).

## Repository Eligibility

To be eligible for a GHAS scan, the repository must have GHAS enabled. A repository can have GHAS enabled if:

1. The repository is public.
2. The repository is private and GHAS has been purchased.
