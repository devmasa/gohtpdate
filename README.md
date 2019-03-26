# gohtpdate

社内でntpが通らないのでMacの時刻が徐々にずれていって困っていた。  
そして、htpdateなるものを知ったのでGo練習がてらCLI作ってみた。  
厳密なhtpdateではない。取ってきて設定するだけ。  
launchctl使うのでOS側で定期実行も可能だけど、練習がてらchannel使ってgoルーチンで定期実行してみる。  
やってみたけどtime.Tickが設定した間隔通りに動かない。でもまあ動くのでよしとする。

## 起動時に実行

Macだと時間設定する時にroot権限必要。  
/Library/LaunchDaemons/配下にplistを配置する。  

```shell
sudo launchctl load xxxxxx.plist
```

## オプション

- i

interval:何分おきに実行するか。そんなにずれないだろうから一日一回でもいいかも。

- p

proxy:プロキシ環境下ならば設定する。ホストとポートを設定する。

- h

host:どのサイトから取得するか。デフォルトはwww.google.co.jp。
