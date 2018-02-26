# Runes

A drop-in replacement for some features of Go's standard `unicode` package.

[![Build Status](https://travis-ci.org/martingallagher/runes.svg)](https://travis-ci.org/martingallagher/runes) [![GoDoc](https://godoc.org/github.com/martingallagher/runes?status.svg)](https://godoc.org/github.com/martingallagher/runes) [![Go Report Card](https://goreportcard.com/badge/github.com/martingallagher/runes)](https://goreportcard.com/report/github.com/martingallagher/runes) [![license](https://img.shields.io/github/license/martingallagher/runes.svg)](https://github.com/martingallagher/runes/blob/master/LICENSE)

This package works by computing an array jump table for fast Unicode property lookup.

Drawbacks include:

- The startup cost of computing the array of properties
- Memory consumption of ~25Mb on x86-64

## Benchmarks

Use `make bench` to run the benchmarks.

- [Intel (x86-64)](#intel)
- [Raspberry Pi 2 Model B (ARMv7)](#raspberry-pi-2)

### Intel

Dell XPS 13 9370 (Intel i7-8550U) / Ubuntu 16.04, CPU in performance mode:

```
benchmark                              old ns/op     new ns/op     delta
BenchmarkIsDigit-8                     34.1          1.10          -96.77%
BenchmarkIsDigitUnsafe-8               33.8          0.55          -98.37%
BenchmarkIsGraphic-8                   127           1.11          -99.13%
BenchmarkIsGraphicUnsafe-8             128           0.83          -99.35%
BenchmarkIsLetter-8                    40.9          1.10          -97.31%
BenchmarkIsLetterUnsafe-8              40.7          0.55          -98.65%
BenchmarkIsLower-8                     30.2          1.12          -96.29%
BenchmarkIsLowerUnsafe-8               30.2          0.83          -97.25%
BenchmarkIsMark-8                      35.0          1.10          -96.86%
BenchmarkIsMarkUnsafe-8                35.1          0.55          -98.43%
BenchmarkIsNumber-8                    30.1          1.10          -96.35%
BenchmarkIsNumberUnsafe-8              30.3          0.83          -97.26%
BenchmarkIsPrint-8                     125           1.10          -99.12%
BenchmarkIsPrintUnsafe-8               125           0.54          -99.57%
BenchmarkIsPunct-8                     30.1          1.11          -96.31%
BenchmarkIsPunctUnsafe-8               29.9          0.82          -97.26%
BenchmarkIsSpace-8                     14.2          1.10          -92.25%
BenchmarkIsSpaceUnsafe-8               14.3          0.55          -96.15%
BenchmarkIsSymbol-8                    30.5          1.11          -96.36%
BenchmarkIsSymbolUnsafe-8              30.2          0.83          -97.25%
BenchmarkIsTitle-8                     14.8          1.10          -92.57%
BenchmarkIsTitleUnsafe-8               14.7          0.55          -96.26%
BenchmarkIsUpper-8                     30.0          1.11          -96.30%
BenchmarkIsUpperUnsafe-8               30.2          0.82          -97.28%
BenchmarkSimpleFold-8                  15.3          0.28          -98.17%
BenchmarkSimpleFoldUnsafe-8            15.1          0.55          -96.36%
BenchmarkToLower-8                     209746        90991         -56.62%
BenchmarkToLowerUnsafe-8               208996        91548         -56.20%
BenchmarkToTitle-8                     222480        101441        -54.40%
BenchmarkToTitleUnsafe-8               219774        102589        -53.32%
BenchmarkToUpper-8                     223247        103880        -53.47%
BenchmarkToUpperUnsafe-8               221264        104025        -52.99%
BenchmarkStringToLowerCase-8           308085        108448        -64.80%
BenchmarkStringToLowerCaseUnsafe-8     308406        91254         -70.41%
BenchmarkStringToTitleCase-8           318338        120817        -62.05%
BenchmarkStringToTitleCaseUnsafe-8     321452        103651        -67.76%
BenchmarkStringToUpperCase-8           325173        120402        -62.97%
BenchmarkStringToUpperCaseUnsafe-8     326598        104489        -68.01%

benchmark                              old MB/s     new MB/s     speedup
BenchmarkIsDigit-8                     234.40       7267.85      31.01x
BenchmarkIsDigitUnsafe-8               236.47       14623.67     61.84x
BenchmarkIsGraphic-8                   62.55        7229.24      115.58x
BenchmarkIsGraphicUnsafe-8             62.04        9668.54      155.84x
BenchmarkIsLetter-8                    195.50       7264.06      37.16x
BenchmarkIsLetterUnsafe-8              196.44       14480.12     73.71x
BenchmarkIsLower-8                     264.99       7117.67      26.86x
BenchmarkIsLowerUnsafe-8               264.74       9621.64      36.34x
BenchmarkIsMark-8                      228.38       7278.42      31.87x
BenchmarkIsMarkUnsafe-8                227.88       14459.79     63.45x
BenchmarkIsNumber-8                    266.12       7285.09      27.38x
BenchmarkIsNumberUnsafe-8              264.07       9664.48      36.60x
BenchmarkIsPrint-8                     63.93        7242.14      113.28x
BenchmarkIsPrintUnsafe-8               63.95        14881.46     232.70x
BenchmarkIsPunct-8                     266.03       7217.83      27.13x
BenchmarkIsPunctUnsafe-8               267.55       9734.01      36.38x
BenchmarkIsSpace-8                     493.22       6351.52      12.88x
BenchmarkIsSpaceUnsafe-8               490.55       12647.13     25.78x
BenchmarkIsSymbol-8                    262.54       7225.87      27.52x
BenchmarkIsSymbolUnsafe-8              265.11       9598.59      36.21x
BenchmarkIsTitle-8                     471.93       6381.03      13.52x
BenchmarkIsTitleUnsafe-8               475.34       12706.11     26.73x
BenchmarkIsUpper-8                     266.55       7236.34      27.15x
BenchmarkIsUpperUnsafe-8               264.68       9711.01      36.69x
BenchmarkSimpleFold-8                  261.92       14454.61     55.19x
BenchmarkSimpleFoldUnsafe-8            264.39       7266.35      27.48x
BenchmarkToLower-8                     76.34        175.98       2.31x
BenchmarkToLowerUnsafe-8               76.62        174.91       2.28x
BenchmarkToTitle-8                     71.97        157.85       2.19x
BenchmarkToTitleUnsafe-8               72.86        156.09       2.14x
BenchmarkToUpper-8                     71.73        154.15       2.15x
BenchmarkToUpperUnsafe-8               72.37        153.93       2.13x
BenchmarkStringToLowerCase-8           51.98        147.66       2.84x
BenchmarkStringToLowerCaseUnsafe-8     51.92        175.48       3.38x
BenchmarkStringToTitleCase-8           50.30        132.54       2.63x
BenchmarkStringToTitleCaseUnsafe-8     49.81        154.49       3.10x
BenchmarkStringToUpperCase-8           49.24        133.00       2.70x
BenchmarkStringToUpperCaseUnsafe-8     49.03        153.25       3.13x
```

### Raspberry Pi 2

Raspberry Pi 2 Model B / Ubuntu 16.04:

```
benchmark                              old ns/op     new ns/op     delta
BenchmarkIsDigit-4                     712           14.7          -97.94%
BenchmarkIsDigitUnsafe-4               716           6.69          -99.07%
BenchmarkIsGraphic-4                   2051          13.4          -99.35%
BenchmarkIsGraphicUnsafe-4             2046          6.72          -99.67%
BenchmarkIsLetter-4                    601           13.4          -97.77%
BenchmarkIsLetterUnsafe-4              600           6.73          -98.88%
BenchmarkIsLower-4                     494           13.4          -97.29%
BenchmarkIsLowerUnsafe-4               495           6.70          -98.65%
BenchmarkIsMark-4                      542           13.4          -97.53%
BenchmarkIsMarkUnsafe-4                542           6.73          -98.76%
BenchmarkIsNumber-4                    492           14.5          -97.05%
BenchmarkIsNumberUnsafe-4              490           6.69          -98.63%
BenchmarkIsPrint-4                     1985          13.4          -99.32%
BenchmarkIsPrintUnsafe-4               1980          6.72          -99.66%
BenchmarkIsPunct-4                     499           13.4          -97.31%
BenchmarkIsPunctUnsafe-4               501           6.72          -98.66%
BenchmarkIsSpace-4                     315           14.4          -95.43%
BenchmarkIsSpaceUnsafe-4               315           6.69          -97.88%
BenchmarkIsSymbol-4                    492           13.4          -97.28%
BenchmarkIsSymbolUnsafe-4              493           6.71          -98.64%
BenchmarkIsTitle-4                     334           13.5          -95.96%
BenchmarkIsTitleUnsafe-4               333           6.72          -97.98%
BenchmarkIsUpper-4                     495           14.7          -97.03%
BenchmarkIsUpperUnsafe-4               495           6.69          -98.65%
BenchmarkSimpleFold-4                  213           4.48          -97.90%
BenchmarkSimpleFoldUnsafe-4            213           4.46          -97.91%
BenchmarkToLower-4                     159897        58494         -63.42%
BenchmarkToLowerUnsafe-4               159290        62179         -60.96%
BenchmarkToTitle-4                     158902        61126         -61.53%
BenchmarkToTitleUnsafe-4               159651        65347         -59.07%
BenchmarkToUpper-4                     159147        61988         -61.05%
BenchmarkToUpperUnsafe-4               159190        55153         -65.35%
BenchmarkStringToLowerCase-4           245836        67834         -72.41%
BenchmarkStringToLowerCaseUnsafe-4     245460        66175         -73.04%
BenchmarkStringToTitleCase-4           245313        81020         -66.97%
BenchmarkStringToTitleCaseUnsafe-4     246914        61785         -74.98%
BenchmarkStringToUpperCase-4           246976        69070         -72.03%
BenchmarkStringToUpperCaseUnsafe-4     246685        55126         -77.65%

benchmark                              old MB/s     new MB/s     speedup
BenchmarkIsDigit-4                     11.24        545.27       48.51x
BenchmarkIsDigitUnsafe-4               11.17        1195.03      106.99x
BenchmarkIsGraphic-4                   3.90         597.51       153.21x
BenchmarkIsGraphicUnsafe-4             3.91         1190.33      304.43x
BenchmarkIsLetter-4                    13.31        595.63       44.75x
BenchmarkIsLetterUnsafe-4              13.33        1188.44      89.16x
BenchmarkIsLower-4                     16.18        595.76       36.82x
BenchmarkIsLowerUnsafe-4               16.16        1194.75      73.93x
BenchmarkIsMark-4                      14.74        595.65       40.41x
BenchmarkIsMarkUnsafe-4                14.74        1188.45      80.63x
BenchmarkIsNumber-4                    16.26        550.34       33.85x
BenchmarkIsNumberUnsafe-4              16.29        1195.30      73.38x
BenchmarkIsPrint-4                     4.03         595.63       147.80x
BenchmarkIsPrintUnsafe-4               4.04         1191.12      294.83x
BenchmarkIsPunct-4                     16.00        595.16       37.20x
BenchmarkIsPunctUnsafe-4               15.96        1191.26      74.64x
BenchmarkIsSpace-4                     22.21        485.16       21.84x
BenchmarkIsSpaceUnsafe-4               22.20        1045.68      47.10x
BenchmarkIsSymbol-4                    16.24        595.77       36.69x
BenchmarkIsSymbolUnsafe-4              16.22        1191.50      73.46x
BenchmarkIsTitle-4                     20.94        519.86       24.83x
BenchmarkIsTitleUnsafe-4               20.97        1042.31      49.70x
BenchmarkIsUpper-4                     16.16        545.92       33.78x
BenchmarkIsUpperUnsafe-4               16.16        1195.17      73.96x
BenchmarkSimpleFold-4                  18.73        893.80       47.72x
BenchmarkSimpleFoldUnsafe-4            18.73        896.38       47.86x
BenchmarkToLower-4                     4.43         12.10        2.73x
BenchmarkToLowerUnsafe-4               4.44         11.39        2.57x
BenchmarkToTitle-4                     4.46         11.58        2.60x
BenchmarkToTitleUnsafe-4               4.43         10.83        2.44x
BenchmarkToUpper-4                     4.45         11.42        2.57x
BenchmarkToUpperUnsafe-4               4.45         12.84        2.89x
BenchmarkStringToLowerCase-4           2.88         10.44        3.62x
BenchmarkStringToLowerCaseUnsafe-4     2.88         10.70        3.72x
BenchmarkStringToTitleCase-4           2.89         8.74         3.02x
BenchmarkStringToTitleCaseUnsafe-4     2.87         11.46        3.99x
BenchmarkStringToUpperCase-4           2.87         10.25        3.57x
BenchmarkStringToUpperCaseUnsafe-4     2.87         12.84        4.47x
```
