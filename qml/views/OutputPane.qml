import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.13

import "../."
import "../controls"

Pane {
    id: root
    padding: 0

    ColumnLayout {
        anchors.fill: parent
        spacing: 0

        OutputHeader {}

        Rectangle {
            Layout.fillWidth: true
            Layout.topMargin: -1
            height: 1
            color: Style.borderColor
        }

        TabBar {
            id: tabbar
            background: Rectangle {
                color: Style.bgColor
            }

            TabButton {
                text: qsTr("Response")
            }

            TabButton {
                text: qsTr("Header/Trailer")
            }
        }

        Rectangle {
            Layout.fillWidth: true
            Layout.topMargin: -1
            Layout.leftMargin: tabbar.width
            height: 1
            color: Style.borderColor
        }

        StackLayout {
            currentIndex: tabbar.currentIndex

            ResponseData {}

            Pane {
                Label { text: "Header content" }
            }
        }

    }
}

