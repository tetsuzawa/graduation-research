# ソフトウェアの製作

## 使用するソフトウェア・プログラムについて

本研究ではシミュレーションとデータ処理のためにPythonとGo言語を主に使用した。

### Pythonについて

Pythonは、汎用のプログラミング言語である。コードがシンプルで扱いやすく設計されており、C言語などに比べて、さまざまなプログラムを分かりやすく、少ないコード行数で書けるといった特徴がある。

標準ライブラリやサードパーティ製のライブラリなど、さまざまな領域に特化した豊富で大規模なツール群が用意され、自らの使用目的に応じて機能を拡張していくことができる。

このような特徴から、様々な試行をしながら行う開発に適している。本研究ではこの特性を活かし、予備実験・データ処理・グラフ作成にPythonを使用している。

### Go言語について

Googleで開発された、汎用のプログラミング言語である。Goは、静的型付け、C言語の伝統に則ったコンパイル言語、メモリ安全性、ガベージコレクション、並行性などの特徴を持つ。また、軽量スレッディングのための機能、Pythonのような動的型付け言語のようなプログラミングの容易性、などの特徴もある。

また、GoはPythonのようにシンプルな記法をもち、誰が書いても同じコードになりやすく学習しやすいため、研究用途に適した言語である。

### プログラミング言語の使い分けについて

Pythonは@で述べた特徴から、様々な試行をしながら行う開発に適している。一方で、速度が遅く、リアルタイム処理には向かない。したがって、本研究ではこの特性を活かし、予備実験・データ処理・グラフ作成にPythonを使用した。

また、Goは言語仕様が現代的で、実行速度が速く、クロスコンパイルが可能といった特徴から、ADFライブラリの作成および実装に使用した。

@TODO

## Pythonを使用した適応フィルタの予備実験

### 音声ファイル処理およびグラフ作成プログラムの制作

本節では後述の実験結果の処理に使用したプログラムの制作について述べる。

音声ファイルの処理には行列演算が必要である。この処理には数値計算を効率的に行うための拡張モジュールであるNumPyを使用した。また、波形のプロットにはNumPyを基盤にしたグラフ描画ライブラリmatplotlibを使用した。
音声ファイルの入出力には標準ライブラリであるwaveや外部ライブラリのPySoundFileを使用した。

@TODO wavファイル

### Python
   - モジュール
      - 波形プロットモジュール
         波形プロット簡易化のモジュールを@に示す。
         @plot_tools.py

      - 音声ファイル入出力モジュール
         音声ファイル入出力のモジュールを@に示す。
         @wave_handler.py

   - デコレータ
      - 関数の実行時間計測用デコレータ
         @decorators.py

      - 関数の情報表示用デコレータ
         @decorators.py

   - ツール
      - 音声ファイル畳み込み用プログラム
         @calc_covolution_wav.py

      - ステレオ音声ファイルをLRモノラル音声に分割するプログラム
         @calc_stereo2monoLR.py

      - 2つのwavファイル間の伝達関数を求めるプログラム
         @calc_pseudo_ir_between_wav_files.py

      - csvファイルの各列のデータをそれぞれwavファイルに変換するプログラム
         @csv_to_wav_each_column.py

      - 目的の音声ファイルからノイズの音声ファイルに含まれる成分を取り除くプログラム（スペクトラムサブトラクション法）
         @calc_subtracted_wav.py

      - 指定秒数分の白色雑音をサンプリング周波数48kHzで生成するプログラム
         @generate_white_noise_as_wav.py

      - csvファイルからmp4またはgif画像を生成するプログラム
         @plot_animation_from_csv.py

      - csvファイルから波形をプロットするプログラム
         @plot_from_csv.py

      - 複数のwavファイルから一枚の図に波形をプロットするプログラム
         @plot_multiwave.py

      - スペクトログラムを表示するプログラム
         @plot_spectrogram_librosa.py


Go言語

- ライブラリ
   後述のツールで使用するライブラリ

- プログラム
   エラーハンドリングはデバッグと簡易化のために`panic`を使用しているが本来はコマンドラインツールには使用するべきではない。

   - 指定した音圧差で音声を合成するプログラム
      - @calc_noise_mix
   - 2つのwavファイルを畳み込むプログラム
      - @convolve_wav
   - 2つのwavファイルを畳み込むプログラム（フーリエ変換を用いた高速版）
      - @convolve_wav_fast
   - 2つのwavファイルもしくはcsvファイルを畳み込むプログラム（フーリエ変換を用いた高速版）
      - @convolve_wav_fast
   - csvファイルのデータからwavファイルを生成するプログラム
      - @csv_to_wav
   - @の実験で用いるADFの設定を記述したJSONファイルを生成するプログラム
      - @drone_json_generator
   - csvファイルのデータから指定したサンプルの平均二乗誤差（MSE）を計算し、csvファイルに新たな列として追記するプログラム
     - @calc_mse_csv
   - PortAudioを使用した多チャンネル収音用プログラム
     - @multirecord
   - DSBファイルからwavファイルに変換するプログラム
     - @wav_to_DSB


@TODO ディレクトリ構造

waveを使用した自作音声入出力モジュールを@に示す。

### 定常なFIRフィルタの動作確認

@TODO


### 既存のADFライブラリについて

pythonで記述された適応信号処理のライブラリとして[Padasip](https://matousc89.github.io/padasip/index.html#padasip)というものが存在する。Padasipは信号のフィルタリング、予測、復元、分類といった適応信号処理を簡易化するために設計されたライブラリである。PadasipではLMS・NLMS、AP、RLSをはじめとする主要なアルゴリズムが一通り実装されている。

## ADFライブラリの制作（Go言語）

### ライブラリの設計

pythonで記述された適応信号処理のライブラリとして[Padasip](https://matousc89.github.io/padasip/index.html#padasip)というものが存在する。Padasipは信号のフィルタリング、予測、復元、分類といった適応信号処理を簡易化するために設計されたライブラリである。PadasipではLMS・NLMS、AP、RLSをはじめとする主要なアルゴリズムが一通り実装されている。

本研究で制作したADFライブラリはPadasipを参考に設計し、Go言語で実装した。

ライブラリのコードを@に示す。

ライブラリの各関数にはユニットテストが書かれており、GitHub Actionsを利用して自動テストが行われるため、一定の質が担保されている。

また、ライセンスはMITを採用しているため、使用、再配布、商用利用などが許可されている。詳しくは[MIT LICENCE](https://raw.githubusercontent.com/tetsuzawa/go-adflib/master/LICENSE)を参照されたい。

### インストール方法

製作したライブラリは[GitHub](https://github.com/tetsuzawa/go-adflib)で公開したため、次のコマンドを実行することでインストールすることができる。

`go get github.com/tetsuzawa/go-adflib/adf`

### 使用方法

以下に基本的なライブラリの使用方法を示す。

- インスタンスの作成
   `NewFilt***`
- フィルタ更新方法
   - リアルタイム処理する場合
      `Adapt`
   - 測定したデータに対して後から処理する場合
      `Run`
- 最適なステップサイズを探す場合
   `ExploreLearning`
- 各種パラメータを取得する場合
   `GetParams`

詳しい使い方は[GoDoc](https://godoc.org/github.com/tetsuzawa/go-adflib)（Goのパッケージリファレンス）を参照されたい。

サンプルプログラムを@に示す。

