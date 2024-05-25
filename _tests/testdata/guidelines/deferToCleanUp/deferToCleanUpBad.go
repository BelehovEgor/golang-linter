package deferToCleanUp

func deferToCleanUpBadEx() int {
	p := deferEx{count: 0}
	p.Lock()
	if p.count < 10 {
		p.Unlock() // want "Use defer to clean up resources such as files and locks."
		return p.count
	}

	p.count++
	newCount := p.count
	p.Unlock() // want "Use defer to clean up resources such as files and locks."

	return newCount
	// easy to miss unlocks due to multiple returns
}
