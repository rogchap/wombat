import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

Dialog {
    id: control 

    property string text

    modal: true
    anchors.centerIn: Overlay.overlay
    width: 500

    title: "Error"

    standardButtons: Dialog.Ok

    header: Label {
        text: control.title
        visible: control.title
        elide: Label.ElideRight
        font.bold: true
        padding: 12
        color: Style.redColor
    }

    background: Rectangle {
        color: Style.bgColor
        border.color: Style.borderColor
    }

    footer: DialogButtonBox {
        visible: count > 0
        background: Rectangle {
            x: 1; y: 1
            width: parent.width - 2
            height: parent.height - 2
            color: Style.bgColor
        }
        delegate: Button {
            hideBorder: false
        }
    }

    TextEdit {
        id: txt
        width: parent.width

        text: control.text
        readOnly: true
        color: Style.textColor
        wrapMode: Text.WordWrap
    }
}
