# eprescription-go
日本の電子処方箋のためのウェブアプリケーションフレームワークおよび実装

2016/04から電子処方箋が解禁されましたが
発表されたガイドライン案だけでは必要な機能について触れられていません。
一応薬剤師をしている私がオープンソースなフレームワークおよび実装を作ることで
どんな機能が必要なのか洗い出したり、議論する場ができると考えています。

電子処方箋サーバは、アクセス数は時間帯の差が激しいと予想され、
自動スケーリングできるGAEがコスト的にいいんじゃないかと。
実行速度やメンテナンス性や抽象化具合、それから私の好みでGo言語を選んでいます。

どこのベンダーもまだ手をつけてないと思いますので、
ぼちぼち書いていきます。




