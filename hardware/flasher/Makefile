default: executable

dependencies: requirements.txt
	pip3 install -r requirements.txt

run: dependencies
	python3 flasher.py

executable: dependencies
	pyinstaller --onefile flasher.py

executable_mac: dependencies
	pyinstaller --onefile --add-binary='/System/Library/Frameworks/Tk.framework/Tk':'tk' --add-binary='/System/Library/Frameworks/Tcl.framework/Tcl':'tcl' flasher.py

zip:
	-rm dist.zip
	cd dist && 7za a -tzip ../dist.zip flasher*
	cd .. && 7za a flasher/dist.zip firmware

clean:
	-rm -r dist.zip dist build
