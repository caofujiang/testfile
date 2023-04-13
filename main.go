package main

/*
func main1() {
	var hKernel32 = syscall.MustLoadDLL("kernel32.dll")
	var hUser32 = syscall.MustLoadDLL("user32.dll")

	// Set up CreateWindowExW parameters for hidden window
	var wcex = syscall.WNDCLASSEX{
		ClassName: syscall.StringToUTF16Ptr("MyHiddenWindowClass"),
		WndProc:   syscall.NewCallback(handleMessage),
	}
	wcex.CbSize = uint32(unsafe.Sizeof(wcex))
	wcex.Style = syscall.CS_HREDRAW | syscall.CS_VREDRAW
	var windowClass = syscall.MustRegisterClassEx(&wcex)

	// Create hidden window
	var hWnd = syscall.CreateWindowEx(
		0,
		syscall.StringToUTF16Ptr("MyHiddenWindowClass"),
		nil,
		0, 0, 0, 0, 0,
		0,
		0,
		windowClass,
		0,
	)

	// Prepare CreateProcess call
	var args []string = os.Args[1:]
	var commandLine []uint16
	for _, arg := range args {
		commandLine = append(commandLine, []rune(arg)...)
		commandLine = append(commandLine, 0)
	}
	var si = new(syscall.StartupInfo)
	var pi = new(syscall.ProcessInformation)
	si.Flags |= syscall.STARTF_USESHOWWINDOW
	si.ShowWindow = syscall.SW_HIDE

	// Create process
	hKernel32.MustFindProc("CreateProcessW").Call(
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(args[0]))),
		uintptr(unsafe.Pointer(&commandLine[0])),
		0,
		false,
		syscall.CREATE_NEW_PROCESS_GROUP|syscall.CREATE_NO_WINDOW,
		0, 0,
		uintptr(unsafe.Pointer(si)),
		uintptr(unsafe.Pointer(pi)),
	)

	// Wait for process to finish (optional)
	hKernel32.MustFindProc("WaitForSingleObject").Call(uintptr(pi.Process), ^uintptr(0))

	// Destroy window
	hUser32.MustFindProc("DestroyWindow").Call(uintptr(hWnd))
}

func handleMessage(hWnd syscall.Handle, msg uint32, wParam, lParam uintptr) uintptr {
	return syscall.DefWindowProc(hWnd, msg, wParam, lParam)
}
*/
