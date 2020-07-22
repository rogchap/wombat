import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.13
import Qt.labs.platform 1.1
import Qt.labs.qmlmodels 1.0

import "views"

ApplicationWindow {
    id: window
    visible: true
    title: "Courier" 
    minimumWidth: 1200
    minimumHeight: 820
    color: Style.bgColor
    flags: Qt.WindowFullscreenButtonHint
   
    // flags: Qt.WindowStaysOnTopHint
    // onActiveFocusItemChanged: print(activeFocusItem)

    header: Header {}

    MainContent {}
}
