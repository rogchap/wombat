import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

ScrollView {
    id: root

    padding: 15

    readonly property var oc: mc.workspaceCtrl.outputCtrl

    clip: true
    
    function css() {
        return `
        <style>
            .name {
                color: ${Style.accentColor2}
            }
            .str {
                color: ${Style.greenColor}
            }
            .bkt {
                color: ${Style.primaryColor}
            }
            .num {
                color: ${Style.purpleColor}
            }
            .bool {
                color: ${Style.accentColor}
            }
        </style>
        `
    }

    TextEdit {
        id: outTxt

        width: root.width - root.padding * 2
        text: css() + oc.output
        color: Style.textColor
        textFormat: TextEdit.RichText
        readOnly: true
        selectByMouse: true
        selectionColor: Style.accentColor2
        wrapMode: TextEdit.Wrap

        MouseArea {
            width: parent.width
            height: parent.height
            cursorShape: Qt.IBeamCursor
            enabled: false
        }

    }

}
