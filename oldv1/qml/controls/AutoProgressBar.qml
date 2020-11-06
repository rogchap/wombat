import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."


ProgressBar {
    id: control

    property bool swap: false

    LayoutMirroring.enabled: swap

    background: Rectangle {
        color: Style.bgColor
        implicitWidth: 200
        implicitHeight: 2
    }

    contentItem: Item {
        implicitWidth: 200
        implicitHeight: 2
            scale: control.mirrored ? -1 : 1

        Rectangle {
            width: control.position * parent.width
            height: parent.height
            radius: 2
            color: Style.primaryColor
        }
    }

    onVisibleChanged: {
        swap = false
        value = 0
    }

    SequentialAnimation {
        loops: Animation.Infinite
        running: control.visible
        PropertyAnimation {
            target: control
            property: "value"
            from: 0
            to: 1
            duration: 3000
            easing.type: Easing.InOutQuart
        }
        PropertyAnimation {
            target: control
            property: "swap"
            to: true
        }
        PropertyAnimation {
            target: control
            property: "value"
            from: 1
            to: 0
            duration: 800
            easing.type: Easing.InOutQuad
        }
        PropertyAnimation {
            target: control
            property: "swap"
            to: false
        }
    }
}

