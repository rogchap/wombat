import QtQuick 2.13
import QtQuick.Layouts 1.13
import QtQuick.Window 2.13

import "views"

Window {
    id: window
    visible: true
    title: "Wombat" 
    minimumWidth: 1200
    minimumHeight: 820
    color: Style.bgColor
    flags: Qt.Window | Qt.WindowFullscreenButtonHint
   
    // flags: Qt.WindowStaysOnTopHint

    MacOSMenuBar {}

    ColumnLayout {
        spacing: 1
        anchors.fill: parent

        Header {
            Layout.fillWidth: true
        }

        MainContent {
            Layout.fillHeight: true
            Layout.fillWidth: true
        }
    }
}
