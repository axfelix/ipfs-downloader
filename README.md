# IPFS Downloader

This is a sample Temporal project that can be used to retry downloading flaky IPFS-over-HTTP URLs 

## Instructions

Ensure you have Go 1.21 or later installed locally, and a local [Temporal Cluster](https://docs.temporal.io/cli) running.

You should be able to view your local cluster's Temporal Web UI at <http://localhost:8233>.

Clone this repository:

```bash
git clone https://github.com/axfelix/ipfs-downloader
```

Run the worker included in the project:

```bash
go run worker/main.go
```

Then, run the starter, and provide both a `--url` and a download `--dir`:

```bash
go run start/main.go --url https://download.library.lol/main/4154000/387eb49ed942fda6a0ef896a98ecceae/%28Literature%20Now%29%20Dan%20Sinykin%20-%20Big%20Fiction_%20How%20Conglomeration%20Changed%20the%20Publishing%20Industry%20and%20American%20Literature-Columbia%20University%20Press%20%282023%29.pdf --dir ~/Desktop
```
