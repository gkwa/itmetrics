# itmetrics

- https://github.com/taylormonacelli/wildcalifornia/blob/master/all.k
- https://raw.githubusercontent.com/taylormonacelli/wildcalifornia/master/all.k

## Install from homebrew

```bash

brew install taylormonacelli/homebrew-tools/itmetrics

```

## Usage

### example

```bash

itmetrics run --manifest=file:///Users/mtm/pdev/taylormonacelli/wildcalifornia/all.k --outdir=trash



```

### example

```bash

rm -rf example*; make && ./itmetrics run --manifest=https://raw.githubusercontent.com/taylormonacelli/wildcalifornia/master/all.k


```

## Build

```bash

make && ./itmetrics run --manifest=all.k

```
