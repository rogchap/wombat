import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "."

CheckBox {
    id: control

    spacing: 10

    contentItem: Label {
        rightPadding: control.indicator.width + control.spacing
        text: control.text
        verticalAlignment: Text.AlignVCenter
    }

    indicator: Rectangle {
        implicitWidth: 24
        implicitHeight: 24
        x: control.width - width - control.rightPadding
        y: control.topPadding + control.availableHeight / 2 - height / 2

        color: Style.bgInputColor
        border.color: Style.borderColor

        Rectangle {
            width: 14
            height: 14
            x: 5
            y: 5
            color: Style.primaryColor
            visible: control.checked
        }
    }
}
