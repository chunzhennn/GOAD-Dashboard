import { baseApi as api } from "./baseApi";
const injectedRtkApi = api.injectEndpoints({
  endpoints: (build) => ({
    getApiPfsenseOpenvpnConnections: build.query<
      GetApiPfsenseOpenvpnConnectionsApiResponse,
      GetApiPfsenseOpenvpnConnectionsApiArg
    >({
      query: () => ({ url: `/api/pfsense/openvpn/connections` }),
    }),
    getApiPveReset: build.query<
      GetApiPveResetApiResponse,
      GetApiPveResetApiArg
    >({
      query: () => ({ url: `/api/pve/reset` }),
    }),
    postApiPveReset: build.mutation<
      PostApiPveResetApiResponse,
      PostApiPveResetApiArg
    >({
      query: () => ({ url: `/api/pve/reset`, method: "POST" }),
    }),
    getApiPveVms: build.query<GetApiPveVmsApiResponse, GetApiPveVmsApiArg>({
      query: () => ({ url: `/api/pve/vms` }),
    }),
    postApiPveVmsReset: build.mutation<
      PostApiPveVmsResetApiResponse,
      PostApiPveVmsResetApiArg
    >({
      query: () => ({ url: `/api/pve/vms/reset`, method: "POST" }),
    }),
    postApiPveVmsStart: build.mutation<
      PostApiPveVmsStartApiResponse,
      PostApiPveVmsStartApiArg
    >({
      query: () => ({ url: `/api/pve/vms/start`, method: "POST" }),
    }),
    postApiPveVmsStop: build.mutation<
      PostApiPveVmsStopApiResponse,
      PostApiPveVmsStopApiArg
    >({
      query: () => ({ url: `/api/pve/vms/stop`, method: "POST" }),
    }),
  }),
  overrideExisting: false,
});
export { injectedRtkApi as api };
export type GetApiPfsenseOpenvpnConnectionsApiResponse =
  /** status 200 OK */ PfsensePfsenseOpenVpnConnection[];
export type GetApiPfsenseOpenvpnConnectionsApiArg = void;
export type GetApiPveResetApiResponse = { last_reset: number };
export type GetApiPveResetApiArg = void;
export type PostApiPveResetApiResponse = Record<string, string>;
export type PostApiPveResetApiArg = void;
export type GetApiPveVmsApiResponse = /** status 200 OK */ ProxmoxVmInfo[];
export type GetApiPveVmsApiArg = void;
export type PostApiPveVmsResetApiResponse = Record<string, string>;
export type PostApiPveVmsResetApiArg = void;
export type PostApiPveVmsStartApiResponse = Record<string, string>;
export type PostApiPveVmsStartApiArg = void;
export type PostApiPveVmsStopApiResponse = Record<string, string>;
export type PostApiPveVmsStopApiArg = void;
export interface PfsensePfsenseOpenVpnConnection {
  common_name?: string;
  connect_time_unix?: number;
  id?: number;
}
export interface ProxmoxVmInfo {
  /** Current CPU usage */
  cpu?: number;
  /** Maximum available CPU count */
  cpus?: number;
  /** Current disk usage */
  disk?: number;
  /** Total disk read (bytes) */
  diskread?: number;
  /** Total disk write (bytes) */
  diskwrite?: number;
  id?: string;
  /** Root disk size (bytes) */
  maxdisk?: number;
  /** Maximum memory (bytes) */
  maxmem?: number;
  /** Current memory usage */
  mem?: number;
  name?: string;
  /** Total network input (bytes) */
  netin?: number;
  /** Total network output (bytes) */
  netout?: number;
  node?: string;
  status?: string;
  /** Running time (seconds) */
  uptime?: number;
}
export const {
  useGetApiPfsenseOpenvpnConnectionsQuery,
  useGetApiPveResetQuery,
  usePostApiPveResetMutation,
  useGetApiPveVmsQuery,
  usePostApiPveVmsResetMutation,
  usePostApiPveVmsStartMutation,
  usePostApiPveVmsStopMutation,
} = injectedRtkApi;
