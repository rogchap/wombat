import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."

Button {
    id: control

    property alias color: text.color

    contentItem: Text {
        id: text

        color: Style.textColor
        text: control.text

        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter
    }

    background: Rectangle {
        implicitWidth: 100
        implicitHeight: 40
        color: control.down ? Style.bgColor2 : Style.bgColor
    }
}
