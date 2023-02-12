Web server that returns 5 most frequently searched words starting with given prefix

Searched words are saved to DB incrementing their frequencies by one. A go routine asynchronously builds the trie from scratch every 5 seconds using the searched words and their frequencies. The trie currently only supports lowercase English letters and it's serialized/deserialized using protobufs.

![alt text](https://github.com/selimabeniacar/autocomplete/blob/main/example.png)
