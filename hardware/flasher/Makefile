dependencies:
	pip3 install -r requirements.txt

run:
	python3 flasher.py

executable:
	pyinstaller --onefile flasher.py

executable_mac:
	pyinstaller --onefile --add-binary='/System/Library/Frameworks/Tk.framework/Tk':'tk' --add-binary='/System/Library/Frameworks/Tcl.framework/Tcl':'tcl' flasher.py

zip:
	-rm dist.zip
	cd dist && zip ../dist.zip flasher*
	cd .. && zip -r flasher/dist.zip firmware