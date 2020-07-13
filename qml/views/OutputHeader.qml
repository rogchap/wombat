import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

Pane {
    id: root

    readonly property var oc: mc.workspaceCtrl.outputCtrl

    implicitHeight: 40
    implicitWidth: parent.width
    
    Status {
        anchors.left: parent.left
        anchors.top: parent.top
        anchors.topMargin: -3
        anchors.leftMargin: -3

        code: oc.status
    }

}
