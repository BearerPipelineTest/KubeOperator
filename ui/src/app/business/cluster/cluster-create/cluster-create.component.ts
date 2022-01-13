import {Component, EventEmitter, OnInit, Output, ViewChild} from '@angular/core';
import {ClusterCreateRequest, CreateNodeRequest} from '../cluster';
import {NgForm} from '@angular/forms';
import {ClusterService} from '../cluster.service';
import {ClrWizard} from '@clr/angular';
import {HostService} from '../../host/host.service';
import {PlanService} from '../../deploy-plan/plan/plan.service';
import {Plan} from '../../deploy-plan/plan/plan';
import {Project} from '../../project/project';
import {ActivatedRoute} from '@angular/router';
import {ManifestService} from '../../manifest/manifest.service';
import {SystemService} from '../../setting/system.service';
import {CommonAlertService} from '../../../layout/common-alert/common-alert.service';
import {AlertLevels} from '../../../layout/common-alert/alert';


@Component({
    selector: 'app-cluster-create',
    templateUrl: './cluster-create.component.html',
    styleUrls: ['./cluster-create.component.css']
})
export class ClusterCreateComponent implements OnInit {

    opened = false;
    item: ClusterCreateRequest = new ClusterCreateRequest();
    options: any = {
        multiple: true,
    };
    hosts: any[] = [];
    masters: any[] = [];
    workers: any[] = [];
    plans: Plan[] = [];
    versions: string[] = [];
    repoList: any[] = [];
    repoCheck = true;
    currentProject: Project;
    nameValid = true;
    nameChecking = false;
    helmVersions: string[] = [];


    clusterCidr = "10.244.0.0/18"
    serviceCidr = "10.244.64.0/18"
    maxPodNum = 110


    podMaxNumOptions = [32, 64, 128, 256];
    serviceMaxNumOptions = [32, 64, 128, 256, 512, 1024, 2048, 4096];


    @ViewChild('wizard', {static: true}) wizard: ClrWizard;
    @ViewChild('basicForm') basicForm: NgForm;
    @ViewChild('seniorForm') seniorForm: NgForm;
    @Output() created = new EventEmitter();

    constructor(private service: ClusterService,
                private hostService: HostService,
                private planService: PlanService,
                private commonAlertService: CommonAlertService,
                private route: ActivatedRoute,
                private settingService: SystemService,
                private manifestService: ManifestService) {
    }


    ngOnInit(): void {
        this.getRegistry();
        this.route.parent.data.subscribe(data => {
            this.currentProject = data.project;
        });

    }

    getRegistry() {
        this.settingService.getRegistry().subscribe(data => {
            this.repoList = (data.items === null) ? [] : data.items
            this.changeArch('amd64');
        });
    }






    reset() {
        this.wizard.reset();
        this.seniorForm.reset();
        this.basicForm.reset();
        this.hosts = [];
        this.masters = [];
        this.workers = [];
        this.versions = [];
        this.nameValid = true;
        this.nameChecking = false;
        this.helmVersions = ['v3', 'v2'];
    }

    setDefaultValue() {
        this.item.provider = 'bareMetal';
        this.item.networkType = 'flannel';
        this.item.ciliumVersion = 'v1.9.5';
        this.item.ciliumNativeRoutingCidr = '10.244.0.0/18';
        this.item.runtimeType = 'docker';
        this.item.dockerStorageDir = '/var/lib/docker';
        this.item.containerdStorageDir = '/var/lib/containerd';
        this.item.flannelBackend = 'vxlan';
        this.item.enableDnsCache = 'disable';
        this.item.dnsCacheVersion = '1.17.0';
        this.item.calicoIpv4poolIpip = 'Always';
        this.item.dockerSubnet = '172.17.0.1/16';
        this.item.kubernetesAudit = 'no';
        this.item.kubeProxyMode = 'iptables';
        this.item.ingressControllerType = 'nginx';
        this.item.projectName = this.currentProject.name;
        this.item.workerAmount = 1;
        this.item.architectures = 'amd64';
        this.item.helmVersion = 'v3';
        this.item.supportGpu = 'disable';
        this.item.yumOperate = 'replace';
        this.item.clusterCidr = "10.244.0.0/18"
        this.item.serviceCidr = "10.244.64.0/18"
        this.item.maxPodNum = 110
    }

    onNameCheck() {
        this.nameChecking = true;
        setTimeout(() => {
            this.service.get(this.item.name).subscribe(data => {
                this.nameValid = false;
                this.nameChecking = false;
            }, error => {
                this.nameChecking = false;
                this.nameValid = true;
            });
        }, 1000);

    }

    open() {
        this.reset();
        this.loadHosts();
        this.loadPlan();
        this.loadVersion();
        this.opened = true;
        this.setDefaultValue();
    }

    onCancel() {
        this.opened = false;
        this.reset();
    }


    toggle(role: string) {
        switch (role) {
            case 'worker':
                const delw = [];
                this.masters.forEach(m => {
                    this.workers.forEach(w => {
                        if (m.id === w.id) {
                            delw.push(w);
                        }
                    });
                });
                const cw = [].concat(this.workers);
                delw.forEach(d => {
                    cw.splice(cw.indexOf(d), 1);
                    this.workers = cw;
                });
                break;
            case 'master':
                const delm = [];
                this.workers.forEach(m => {
                    this.masters.forEach(w => {
                        if (m.id === w.id) {
                            delm.push(w);
                        }
                    });
                });
                const cm = [].concat(this.masters);
                delm.forEach(d => {
                    cm.splice(cm.indexOf(d), 1);
                    this.masters = cm;
                });
                break;
        }
    }

    loadHosts() {
        this.hostService.listByProjectName(this.currentProject.name).subscribe(data => {
            const list = [];
            data.items.filter((host) => {
                return host.status === 'Running';

            }).forEach(h => {
                if (!h.clusterId) {
                    list.push({id: h.name, text: h.name, disabled: false});
                }
            });
            this.hosts = list;
        });
    }

    loadPlan() {
        this.planService.listByProjectName(this.currentProject.name).subscribe(data => {
            this.plans = data.items;
        });
    }

    getDefaultTunnelMode() {
        if (this.item.flannelBackend === 'Overlay') {
            this.item.ciliumTunnelMode = 'vxlan'
        } else {
            this.item.ciliumTunnelMode = 'disabled'
        }
    }

    getDefaultFlannelBackend() {
        if (this.item.networkType === 'cilium') {
            this.item.flannelBackend = 'Overlay'
            this.item.ciliumTunnelMode = 'vxlan'
        } else {
            this.item.flannelBackend = 'vxlan'
        }
    }

    loadVersion() {
        this.manifestService.listActive().subscribe(data => {
            for (const m of data) {
                this.versions.push(m.name);
            }
            this.item.version = data[0].name;
        });
    }

    fullNodes() {
        this.item.nodes = [];
        this.masters.forEach(m => {
            const node = new CreateNodeRequest();
            node.hostName = m.id;
            node.role = 'master';
            this.item.nodes.push(node);
        });
        this.workers.forEach(m => {
            const node = new CreateNodeRequest();
            node.hostName = m.id;
            node.role = 'worker';
            this.item.nodes.push(node);
        });
    }

    onSubmit() {
        if (this.item.ciliumTunnelMode === 'flannelBackend') {
            this.item.ciliumTunnelMode = 'disable'
        }
        this.service.create(this.item).subscribe(data => {
            this.opened = false;
            this.created.emit();
        }, error => {
            this.opened = false;
            this.commonAlertService.showAlert(error.error.msg, AlertLevels.ERROR);
        });
    }

    changeArch(type) {
        this.repoCheck = true;
        let isAmdExit: boolean = false;
        let isArmExit: boolean = false;
        switch (type) {
            case 'amd64':
                for (const repo of this.repoList) {
                    if (repo.architecture === 'x86_64') {
                        isAmdExit = true;
                        break;
                    }
                }
                this.repoCheck = isAmdExit;
                break;
            case 'arm64':
                for (const repo of this.repoList) {
                    if (repo.architecture === 'aarch64') {
                        isArmExit = true;
                        break;
                    }
                }
                this.repoCheck = isArmExit;
                break;
            case 'all':
                for (const repo of this.repoList) {
                    if (repo.architecture === 'x86_64') {
                        isAmdExit = true;
                        continue;
                    }
                    if (repo.architecture === 'aarch64') {
                        isArmExit = true;
                        continue;
                    }
                }
                this.repoCheck = isAmdExit && isArmExit;
                break;
        }
        if (type !== 'amd64') {
            this.item.helmVersion = 'v3';
            this.helmVersions = ['v3'];
        } else {
            this.helmVersions = ['v3', 'v2'];
        }
    }

    getHostName(hosts: any) {
        let hostName = '';
        for (const h of hosts) {
            hostName = h['text'] + ',' + hostName;
        }
        return hostName;
    }



}
