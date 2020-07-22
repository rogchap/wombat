import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.13
import QtQuick.Dialogs 1.1

import "../."
import "../controls"

Modal {
    id: root

    readonly property var wc: mc.workspaceCtrl

    headerText: qsTr("Workspace")

    ColumnLayout {
        spacing: 0

        TabBar {
            id: tabbar
            background: Rectangle {
                color: Style.bgColor
            }

            TabButton {
                text: qsTr("Basic")
            }

            TabButton {
                text: qsTr("TLS")
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

            WorkspaceOptionsBasic {
                id: basics
            }

            WorkspaceOptionsTls {}

        }

        Rectangle {
            height: 1
            Layout.fillWidth: true
            Layout.topMargin: 10
            Layout.bottomMargin: 10
            color: Style.borderColor
        }

        Button {
            Layout.alignment: Qt.AlignRight
            bgColor: Style.accentColor3

            text: qsTr("Connect")

            onClicked: {
                let err = wc.processProtos()
                if (err) {
                    print(err)
                    return
                }
                err = wc.connect(basics.addr)
                if (err) {
                    print(err)
                    return
                }
                root.close()
            }
        }
    }



}
