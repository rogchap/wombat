import QtQuick 2.13
import QtQuick.Window 2.13
import QtQuick.Layouts 1.13

import "../controls"

Window {
    minimumWidth: 250
    minimumHeight: 170
    maximumWidth: minimumWidth
    maximumHeight: minimumHeight

    modality: Qt.ApplicationModal
    flags: Qt.Dialog

    Pane {
        anchors.fill: parent

        ColumnLayout {
            width: parent.width

            spacing: 5

            Image {
                width: 80
                height: width
                Layout.alignment: Qt.AlignCenter
                Layout.preferredWidth: width
                Layout.preferredHeight: height
                clip: true
                source: "../img/icon_128x128@2x.png"
            }

            Label {
                Layout.alignment: Qt.AlignCenter
                text: "Wombat"
                font.weight: Font.DemiBold
            }

            Label {
                Layout.alignment: Qt.AlignCenter
                text: "Version " + mc.version
                font.pointSize: 12
            }

            Label {
                Layout.alignment: Qt.AlignCenter
                text: "Copyright 2020, Roger Chapman"
                font.pointSize: 12
            }
        }
    }
}
