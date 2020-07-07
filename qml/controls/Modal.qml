import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "."

Popup {
    id: control

    property alias headerText: lblHeader.text

    anchors.centerIn: Overlay.overlay

    margins: 40
    topPadding: header.height + 15

    modal: true
    focus: true
    clip: true
    closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent

    background: Rectangle {

        color: Style.bgColor
        border.color: Style.borderColor

        Pane {
            id: header
            x: 1
            y: 1

            width: control.width - 2
            height: 40

            Label {
                id: lblHeader
                font.weight: Font.DemiBold
            }

            AbstractButton {
                text: "×"
                font {
                    weight: Font.DemiBold
                    pointSize: 22
                }
                anchors {
                    right: parent.right
                    top: parent.top
                    topMargin: -8
                }

                contentItem: Label {
                    text: parent.text
                }

                onClicked: control.close()
            }
        }

        Rectangle {
            width: control.width
            height: 1
            color: Style.borderColor
            anchors.bottom: header.bottom
        }

    }

}

