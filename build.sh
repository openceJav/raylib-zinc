#!/bin/bash

echo 'Generating app icon..'
ICONFILE='assets/icons/icon_1024x1024.png'
ICONSET='assets/icons/icons.iconset'
mkdir $ICONSET
sips -z 16 16     $ICONFILE --out $ICONSET/icon_16x16.png
sips -z 32 32     $ICONFILE --out $ICONSET/icon_16x16@2x.png
sips -z 32 32     $ICONFILE --out $ICONSET/icon_32x32.png
sips -z 64 64     $ICONFILE --out $ICONSET/icon_32x32@2x.png
sips -z 128 128   $ICONFILE --out $ICONSET/icon_128x128.png
sips -z 256 256   $ICONFILE --out $ICONSET/icon_128x128@2x.png
sips -z 256 256   $ICONFILE --out $ICONSET/icon_256x256.png
sips -z 512 512   $ICONFILE --out $ICONSET/icon_256x256@2x.png
sips -z 512 512   $ICONFILE --out $ICONSET/icon_512x512.png
cp $ICONFILE $ICONSET/icon_512x512@2x.png
iconutil -c icns $ICONSET
rm -R $ICONSET

