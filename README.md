# gparms
Gets urls from stdin and outputs the urls with parameters in them.

**Usage Full**
```
cat allSubs.txt | gparms | qsreplace | nuceli -t ~/nuceli-templates/vulnerabilities/generic
```

**Install**
```
pip install gparms==0.0.8
```

**Usage 1**
```
cat allUrls.txt | gparms
```

**Before**
```
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=4
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=5
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=6
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=7
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=8
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=9
http://example.com/location/index.ch2
http://example.com/location/printableindex.ch2?selectedCounter=1
http://example.com/robots.txt
http://example.com/specification/index.ch2
http://example.com/specification/printableindex.ch2
http://example.com/upload/buildinglogoimage/161-44logonew.gif
http://example.com/upload/buildingmodulelinkimageunselected/1543-navhcontact.gif
http://example.com/upload/buildingmodulelinkimageunselected/1547-navhfeatures.gif
```

**After**
```
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=4
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=5
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=6
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=7
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=8
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=9
http://example.com/location/printableindex.ch2?selectedCounter=1
```

**Usage 2**
```
cat allUrls.txt | gparms -s
```

**Before**
```
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=4&uid=123&count=test
```

**After**
```
http://example.com/leasing/suiteplanpopup.ch2?selectedCounter=4
http://example.com/leasing/suiteplanpopup.ch2?uid=123
http://example.com/leasing/suiteplanpopup.ch2?count=test
```
