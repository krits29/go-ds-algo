package main

type Lock struct{

	isLocked bool
	lockedBy <-chan
	lockCount int
}

func lock() error{
	for isLocked && currentThread != lockedBy {

		wait()
	}
	isLocked = true
	lockedBy = chan
	lockCount++

}

func unlock() {
	// if lock is not locked or current chan is not the owner
	if !isLocked || currentThread != lockedBy {
		return
	}

	lockCount--
	if lockCount == 0 {
		isLocked = false
		notify
	}
}