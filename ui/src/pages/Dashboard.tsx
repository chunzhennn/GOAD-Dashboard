import { useEffect, useState } from 'react';
import { 
  useGetApiPveVmsQuery, 
  usePostApiPveVmsStartMutation, 
  usePostApiPveVmsStopMutation, 
  usePostApiPveVmsResetMutation,
  useGetApiPfsenseOpenvpnConnectionsQuery,
  useGetApiPveResetQuery,
  usePostApiPveResetMutation,
  ProxmoxVmInfo,
  PfsensePfsenseOpenVpnConnection
} from '../store/api';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "../components/ui/card";
import { Button } from "../components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "../components/ui/tabs";
import { Badge } from "../components/ui/badge";
import { Separator } from "../components/ui/separator";
import { ScrollArea } from "../components/ui/scroll-area";
import { RefreshCw, Play, Square, Server, Activity, UserCheck, RotateCw, AlertTriangle } from 'lucide-react';
import { formatDistance } from 'date-fns';
import { cn } from '@/lib/utils';

export function Dashboard() {
  // Get VM list
  const { 
    data: vms = [], 
    isLoading: isLoadingVms, 
    isFetching: isFetchingVms,
    refetch: refetchVms 
  } = useGetApiPveVmsQuery();
  
  // Get VPN connections
  const { 
    data: vpnConnections = [], 
    isLoading: isLoadingVpnConnections,
    isFetching: isFetchingVpnConnections,
    refetch: refetchVpnConnections
  } = useGetApiPfsenseOpenvpnConnectionsQuery();
  
  // Get reset history
  const { 
    data: resetHistory, 
    isLoading: isLoadingResetHistory,
    isFetching: isFetchingResetHistory,
    refetch: refetchResetHistory
  } = useGetApiPveResetQuery();

  // VM operation mutations
  const [startVM] = usePostApiPveVmsStartMutation();
  const [stopVM] = usePostApiPveVmsStopMutation();
  const [resetVM] = usePostApiPveVmsResetMutation();
  const [restoreSnapshots, { isLoading: isRestoringSnapshots }] = usePostApiPveResetMutation();

  // Refresh interval (3 seconds)
  const [refreshInterval, setRefreshInterval] = useState(3000);
  
  // Refresh data periodically
  useEffect(() => {
    const intervalId = setInterval(() => {
      refetchVms();
      refetchVpnConnections();
      refetchResetHistory();
    }, refreshInterval);
    
    return () => { clearInterval(intervalId); };
  }, [refetchVms, refetchVpnConnections, refetchResetHistory, refreshInterval]);

  // Handle VM operations
  const handleStartAllVMs = async () => {
    try {
      await startVM();
      refetchVms();
    } catch (error) {
      console.error('Failed to start all VMs:', error);
    }
  };

  const handleStopAllVMs = async () => {
    try {
      await stopVM();
      refetchVms();
    } catch (error) {
      console.error('Failed to stop all VMs:', error);
    }
  };

  const handleResetAllVMs = async () => {
    try {
      await resetVM();
      refetchVms();
    } catch (error) {
      console.error('Failed to reset all VMs:', error);
    }
  };

  const handleRestoreSnapshots = async () => {
    try {
      await restoreSnapshots();
      refetchVms();
      refetchResetHistory();
    } catch (error) {
      console.error('Failed to restore snapshots:', error);
    }
  };

  // Group VMs by status
  const runningVMs = vms.filter(vm => vm.status === 'running');
  const stoppedVMs = vms.filter(vm => vm.status !== 'running');

  const isLoading = isLoadingVms || isLoadingVpnConnections || isLoadingResetHistory;
  const isFetching = isFetchingVms || isFetchingVpnConnections || isFetchingResetHistory;

  // Format bytes
  const formatBytes = (bytes: number | undefined) => {
    if (bytes === undefined) return 'N/A';
    
    const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
    if (bytes === 0) return '0 B';
    
    const i = Math.floor(Math.log(bytes) / Math.log(1024));
    return parseFloat((bytes / Math.pow(1024, i)).toFixed(2)) + ' ' + sizes[i];
  };

  // Format timestamp
  const formatTimestamp = (timestamp: number) => {
    return new Date(timestamp * 1000).toLocaleString();
  };

  // Get last reset time
  const getLastResetTime = () => {
    if (resetHistory?.last_reset === undefined || resetHistory.last_reset === 0) {
      return 'Never';
    }
    return formatTimestamp(resetHistory.last_reset);
  };

  return (
    <div className="container px-4 md:px-6 py-6 max-w-7xl">
      {/* Status indicator */}
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-3xl font-bold tracking-tight">GOAD Dashboard</h1>
        <div className="flex items-center gap-4">
          {isLoading ? (
            <div className="flex items-center">
              <RefreshCw className="h-4 w-4 animate-spin mr-2" />
              <span className="text-sm text-muted-foreground">Refreshing...</span>
            </div>
          ) : (
            <Button 
              variant="outline" 
              size="sm" 
              onClick={() => {
                refetchVms();
                refetchVpnConnections();
                refetchResetHistory();
              }}
            >
              <RefreshCw className={cn("h-4 w-4 mr-2", isFetching ? "animate-spin" : "")} />
              Refresh
            </Button>
          )}
        </div>
      </div>

      {/* Status overview */}
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 mb-8">
        <Card>
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-medium flex items-center">
              <Server className="h-4 w-4 mr-2" />
              Running VMs
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{runningVMs.length} / {vms.length}</div>
          </CardContent>
        </Card>
        
        <Card>
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-medium flex items-center">
              <Activity className="h-4 w-4 mr-2" />
              Last Reset Time of GOAD
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{getLastResetTime()}</div>
          </CardContent>
        </Card>
        
        <Card>
          <CardHeader className="pb-2">
            <CardTitle className="text-sm font-medium flex items-center">
              <UserCheck className="h-4 w-4 mr-2" />
              Active VPN Connections
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold text-emerald-500">{vpnConnections.length}</div>
          </CardContent>
        </Card>
      </div>

      {/* Main content tabs */}
      <Tabs defaultValue="vms" className="mb-8">
        <TabsList className="mb-4">
          <TabsTrigger value="vms">Virtual Machines</TabsTrigger>
          <TabsTrigger value="connections">VPN Connections</TabsTrigger>
        </TabsList>
        
        <TabsContent value="vms">
          <div className="space-y-6">
            {/* All VMs */}
            <div>
              <div className="flex items-center justify-between mb-4">
                <h2 className="text-xl font-semibold flex items-center">
                  Virtual Machines
                </h2>
                <div className="flex items-center gap-2">
                  <Button 
                    variant="default" 
                    size="sm"
                    onClick={handleStartAllVMs}
                    disabled={runningVMs.length === vms.length}
                  >
                    <Play className="h-4 w-4 mr-2" />
                    Start
                  </Button>
                  <Button 
                    variant="destructive" 
                    size="sm"
                    onClick={handleStopAllVMs}
                    disabled={stoppedVMs.length === vms.length}
                  >
                    <Square className="h-4 w-4 mr-2" />
                    Stop
                  </Button>
                  <Button 
                    variant="secondary" 
                    size="sm"
                    onClick={handleResetAllVMs}
                    disabled={vms.length === 0}
                  >
                    <RotateCw className="h-4 w-4 mr-2" />
                    Reset
                  </Button>
                  <Button 
                    variant="outline" 
                    size="sm"
                    onClick={handleRestoreSnapshots}
                    disabled={isRestoringSnapshots || vms.length === 0}
                  >
                    <RotateCw className="h-4 w-4 mr-2" />
                    Restore
                  </Button>
                </div>
              </div>
              
              {vms.length === 0 ? (
                <Card>
                  <CardContent className="flex items-center justify-center py-6">
                    <p className="text-muted-foreground flex items-center">
                      <AlertTriangle className="h-4 w-4 mr-2" />
                      No VMs found
                    </p>
                  </CardContent>
                </Card>
              ) : (
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                  {[...vms].sort((a, b) => a.id!.localeCompare(b.id!)).map((vm: ProxmoxVmInfo) => {
                    const isRunning = vm.status === 'running';
                    
                    return (
                      <Card key={vm.id} className={isRunning ? "border-blue-300" : "border-red-300"}>
                        <CardHeader className="pb-2">
                          <div className="flex justify-between items-center">
                            <CardTitle>{vm.name || 'Unknown VM'}</CardTitle>
                            <Badge 
                              variant="outline" 
                              className={isRunning 
                                ? "bg-blue-50 text-blue-700 border-blue-200" 
                                : "bg-red-50 text-red-700 border-red-200"}
                            >
                              {vm.status}
                            </Badge>
                          </div>
                          <CardDescription>ID: {vm.id} • Node: {vm.node}</CardDescription>
                        </CardHeader>
                        <CardContent>
                          {isRunning ? (
                            <>
                              <div className="grid grid-cols-2 gap-4 text-sm">
                                <div>
                                  <div className="text-muted-foreground mb-1">CPU</div>
                                  <div className="font-medium">
                                    {vm.cpu !== undefined ? `${(vm.cpu * 100).toFixed(1)}%` : 'N/A'} 
                                    {vm.cpus && ` (${vm.cpus} cores)`}
                                  </div>
                                  <div className="w-full h-2 bg-blue-100 rounded-full mt-1">
                                    <div 
                                      className="h-full bg-blue-500 rounded-full" 
                                      style={{ width: `${vm.cpu ? vm.cpu * 100 : 0}%` }}
                                    ></div>
                                  </div>
                                </div>
                                <div>
                                  <div className="text-muted-foreground mb-1">Memory</div>
                                  <div className="font-medium">
                                    {formatBytes(vm.mem)} / {formatBytes(vm.maxmem)}
                                  </div>
                                  <div className="w-full h-2 bg-blue-100 rounded-full mt-1">
                                    <div 
                                      className="h-full bg-blue-500 rounded-full" 
                                      style={{ width: `${vm.maxmem && vm.mem ? (vm.mem / vm.maxmem) * 100 : 0}%` }}
                                    ></div>
                                  </div>
                                </div>
                              </div>
                              
                              <div className="grid grid-cols-2 gap-4 text-sm mt-4">
                                <div>
                                  <div className="text-muted-foreground mb-1">Network In</div>
                                  <div className="font-medium">{formatBytes(vm.netin)}</div>
                                </div>
                                <div>
                                  <div className="text-muted-foreground mb-1">Network Out</div>
                                  <div className="font-medium">{formatBytes(vm.netout)}</div>
                                </div>
                              </div>
                            </>
                          ) : (
                            <div className="grid grid-cols-2 gap-4 text-sm">
                              <div>
                                <div className="text-muted-foreground mb-1">CPU</div>
                                <div className="font-medium">
                                  {vm.cpus ? `${vm.cpus} cores` : 'N/A'}
                                </div>
                              </div>
                              <div>
                                <div className="text-muted-foreground mb-1">Memory</div>
                                <div className="font-medium">{formatBytes(vm.maxmem)}</div>
                              </div>
                            </div>
                          )}
                        </CardContent>
                      </Card>
                    );
                  })}
                </div>
              )}
            </div>
          </div>
        </TabsContent>
        
        <TabsContent value="connections">
          <Card>
            <CardHeader>
              <CardTitle className="flex justify-between items-center">
                <div className="flex items-center">
                  <UserCheck className="h-5 w-5 mr-2" />
                  OpenVPN Connections
                </div>
                <Badge variant="outline">{vpnConnections.length} active</Badge>
              </CardTitle>
            </CardHeader>
            <CardContent>
              {vpnConnections.length === 0 ? (
                <div className="text-center py-6 text-muted-foreground">
                  No active connections
                </div>
              ) : (
                <ScrollArea className="h-[300px]">
                  <div className="space-y-2">
                    <div className="grid grid-cols-12 text-xs font-medium text-muted-foreground">
                      <div className="col-span-1">#</div>
                      <div className="col-span-5">User</div>
                      <div className="col-span-6">Connection Time</div>
                    </div>
                    <Separator />
                    <div className="space-y-2">
                      {vpnConnections.map((connection: PfsensePfsenseOpenVpnConnection) => (
                        <div 
                          key={connection.id} 
                          className="grid grid-cols-12 py-2"
                        >
                          <div className="col-span-1 font-medium">{connection.id}</div>
                          <div className="col-span-5 font-medium">{connection.common_name || 'Unknown'}</div>
                          <div className="col-span-6">
                            {connection.connect_time_unix ? formatDistance(connection.connect_time_unix * 1000, new Date(), {addSuffix: true}) : 'Unknown'}
                          </div>
                        </div>
                      ))}
                    </div>
                  </div>
                </ScrollArea>
              )}
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  );
} 