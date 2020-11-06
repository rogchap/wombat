import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."

AbstractButton {
    id: control

    property color color

    focusPolicy: Qt.NoFocus

    height: 17
    width: 17

    Label {
        id: label

        anchors.top: control.top
        anchors.left: control.left
        anchors.topMargin: -6
        anchors.leftMargin: 1

        
        color: control.color
        text: "×"
        font {
            weight: Font.DemiBold
            pointSize: 22
        }
    }

    background: Rectangle {
        rotation: control.rotation - 90
        color: control.down ? Style.bgColor2 : Style.bgColor
    }
}
