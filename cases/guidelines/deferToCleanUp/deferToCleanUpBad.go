package deferToCleanUp

func deferToCleanUpBadEx() int {
	p := deferEx{count: 0}
	p.Lock()
	if p.count < 10 {
		p.Unlock()
		return p.count
	}

	p.count++
	newCount := p.count
	p.Unlock()

	return newCount
	// easy to miss unlocks due to multiple returns
}
