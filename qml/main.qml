import QtQuick 2.13
import QtQuick.Controls 2.13

import "views"

ApplicationWindow {
    id: window
    visible: true
    title: "Wombat" 
    minimumWidth: 1200
    minimumHeight: 820
    color: Style.bgColor
    flags: Qt.WindowFullscreenButtonHint
   
    // flags: Qt.WindowStaysOnTopHint

    MacOSMenuBar {}

    header: Header {}

    MainContent {}
}
