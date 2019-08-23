yml ノート
ansible 書き方

始まりは慣習的に

`---`

変数の使用方法

ダブルクォーテーション`""`で囲ったのちに `{{ }}`で囲う


` "{{ hoge_hoge }}" `

コマンド出力の取得方法

```
---
- name: sample 
  shell: echo hogehoge
  register: echo_res

- name: debug echo_res
  debug:
    msg: "{{ echo_res.stdout }}"
```
`.stdout`に標準出力が入っている

