SITE_PACKAGES := $(shell pip show pip | grep '^Location' | cut -f2 -d':')

default: executable

$(SITE_PACKAGES): requirements.txt
	pip3 install -r requirements.txt

run: $(SITE_PACKAGES)
	python3 flasher.py

executable: $(SITE_PACKAGES)
	pyinstaller --onefile flasher.py

executable_mac: $(SITE_PACKAGES)
	pyinstaller --onefile --add-binary='/System/Library/Frameworks/Tk.framework/Tk':'tk' --add-binary='/System/Library/Frameworks/Tcl.framework/Tcl':'tcl' flasher.py

zip:
	-rm dist.zip
	cd dist && 7za a -tzip ../dist.zip flasher*
	cd .. && 7za a flasher/dist.zip firmware

clean:
	-rm -r dist.zip dist build
