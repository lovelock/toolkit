package str

//ChunkByLength 按长度把字符串切割成切片
func ChunkByLength(s string, n int) []string {
	if len(s) < n {
		return []string{s}
	}

	numOfChunks := len(s) / n
	if n*numOfChunks < len(s) {
		numOfChunks++
	}

	var chunks []string
	for i := 0; i < numOfChunks; i++ {
		var chunk string
		if i != numOfChunks-1 {
			chunk = s[n*i : n*(i+1)]
		} else {
			chunk = s[n*i:]
		}

		chunks = append(chunks, chunk)
	}

	return chunks
}
