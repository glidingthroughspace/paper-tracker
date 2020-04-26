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
	cd dist && zip ../dist.zip flasher*
	cd .. && zip -r flasher/dist.zip firmware
