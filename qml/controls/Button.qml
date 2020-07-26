import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."

Button {
    id: control

    property alias color: text.color
    property color bgColor: Style.bgColor
    property bool hideBorder: true

    focusPolicy: Qt.NoFocus

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
        color: control.down ? Style.bgColor2 : control.bgColor
        border {
            color: Style.borderColor
            width: control.hideBorder ? 0 : 1
        }
    }
}
