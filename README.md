TRPG libray
====


## Descriotion
　TRPGのオンラインセッションのためのキャラクターシート保管庫です。  
主に神我狩のために作られています。  

## Requirement

このソフトウェアはGoogleが開発しているプログラミング言語Goによって作られています  
go(golang) 1.5 or later  

また、ユーザーが使用するサーバー(さくらのレンタルサーバーやVPS、ロリポップなど)のOSに  
よって動作できる実行ファイルが異なります。  
将来的には、こちらで実行形式のファイルを用意するのでそちらをご利用ください。

## Install

　DocumentRoot以下cgiを動かせる場所に置いて  
  
$go build -t index.cgi  
とshellに入力してください  
場合によってはpublic/storeの権限を変更してください  
$chmod 755 public/store


## LICENSE
MITL 
see LICENSE
## AUTHOR

[keizo](https://github.com/KeizoBookman)


## Contact me
twitter : @keizo_bookman  
mail: keizo.bookman at gmail.com  
