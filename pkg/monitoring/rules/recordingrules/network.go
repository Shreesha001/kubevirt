/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright The KubeVirt Authors.
 */

package recordingrules

import (
	"github.com/rhobs/operator-observability-toolkit/pkg/operatormetrics"
	"github.com/rhobs/operator-observability-toolkit/pkg/operatorrules"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// networkRecordingRules pre-aggregates the raw per-interface VMI network counters
// (kubevirt_vmi_network_*_total) into per-VMI rate vectors.  Pre-computing these
// at scrape time keeps alert and dashboard PromQL expressions simple and avoids
// re-evaluating the same rate() window for every consumer.
//
// Naming follows the KubeVirt convention: <level>:kubevirt_<metric>:<operation>
//   - "vmi" level  – aggregated by (name, namespace, interface)
//   - "cluster" level – summed across the whole cluster
var networkRecordingRules = []operatorrules.RecordingRule{
	// ── Per-VMI throughput (bytes/s) ───────────────────────────────────────────
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "vmi:kubevirt_vmi_network_receive_bytes:rate5m",
			Help: "Per-VMI network receive throughput in bytes per second, averaged over 5 minutes.",
		},
		MetricType: operatormetrics.GaugeType,
		Expr: intstr.FromString(
			"sum by (name, namespace, interface) (rate(kubevirt_vmi_network_receive_bytes_total[5m]))",
		),
	},
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "vmi:kubevirt_vmi_network_transmit_bytes:rate5m",
			Help: "Per-VMI network transmit throughput in bytes per second, averaged over 5 minutes.",
		},
		MetricType: operatormetrics.GaugeType,
		Expr: intstr.FromString(
			"sum by (name, namespace, interface) (rate(kubevirt_vmi_network_transmit_bytes_total[5m]))",
		),
	},

	// ── Per-VMI packet rates ───────────────────────────────────────────────────
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "vmi:kubevirt_vmi_network_receive_packets:rate5m",
			Help: "Per-VMI network receive packet rate per second, averaged over 5 minutes.",
		},
		MetricType: operatormetrics.GaugeType,
		Expr: intstr.FromString(
			"sum by (name, namespace, interface) (rate(kubevirt_vmi_network_receive_packets_total[5m]))",
		),
	},
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "vmi:kubevirt_vmi_network_transmit_packets:rate5m",
			Help: "Per-VMI network transmit packet rate per second, averaged over 5 minutes.",
		},
		MetricType: operatormetrics.GaugeType,
		Expr: intstr.FromString(
			"sum by (name, namespace, interface) (rate(kubevirt_vmi_network_transmit_packets_total[5m]))",
		),
	},

	// ── Per-VMI error rates ────────────────────────────────────────────────────
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "vmi:kubevirt_vmi_network_receive_errors:rate5m",
			Help: "Per-VMI network receive error rate per second, averaged over 5 minutes.",
		},
		MetricType: operatormetrics.GaugeType,
		Expr: intstr.FromString(
			"sum by (name, namespace, interface) (rate(kubevirt_vmi_network_receive_errors_total[5m]))",
		),
	},
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "vmi:kubevirt_vmi_network_transmit_errors:rate5m",
			Help: "Per-VMI network transmit error rate per second, averaged over 5 minutes.",
		},
		MetricType: operatormetrics.GaugeType,
		Expr: intstr.FromString(
			"sum by (name, namespace, interface) (rate(kubevirt_vmi_network_transmit_errors_total[5m]))",
		),
	},

	// ── Per-VMI packet-drop rates ─────────────────────────────────────────────
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "vmi:kubevirt_vmi_network_receive_packets_dropped:rate5m",
			Help: "Per-VMI network receive packet drop rate per second, averaged over 5 minutes.",
		},
		MetricType: operatormetrics.GaugeType,
		Expr: intstr.FromString(
			"sum by (name, namespace, interface) (rate(kubevirt_vmi_network_receive_packets_dropped_total[5m]))",
		),
	},
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "vmi:kubevirt_vmi_network_transmit_packets_dropped:rate5m",
			Help: "Per-VMI network transmit packet drop rate per second, averaged over 5 minutes.",
		},
		MetricType: operatormetrics.GaugeType,
		Expr: intstr.FromString(
			"sum by (name, namespace, interface) (rate(kubevirt_vmi_network_transmit_packets_dropped_total[5m]))",
		),
	},

	// ── Cluster-wide throughput (useful for capacity dashboards) ──────────────
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "cluster:kubevirt_vmi_network_receive_bytes:sum",
			Help: "Total cluster-wide VMI network receive throughput in bytes per second, averaged over 5 minutes.",
		},
		MetricType: operatormetrics.GaugeType,
		Expr:       intstr.FromString("sum(rate(kubevirt_vmi_network_receive_bytes_total[5m]))"),
	},
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "cluster:kubevirt_vmi_network_transmit_bytes:sum",
			Help: "Total cluster-wide VMI network transmit throughput in bytes per second, averaged over 5 minutes.",
		},
		MetricType: operatormetrics.GaugeType,
		Expr:       intstr.FromString("sum(rate(kubevirt_vmi_network_transmit_bytes_total[5m]))"),
	},
}
