<template>
  <div class="image-container">
    <div id="lineChart" :style="{ width: '600px', height: '600px' }" />
    <div id="myChart" :style="{ width: '600px', height: '600px' }" />
  </div>
</template>

<script>
import { getChart } from '@/api/chart'

export default {
  name: 'Hello',
  data() {
    return {
      chartData: [] // 添加一个变量用于存储图表数据
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      getChart()
        .then(response => {
          this.chartData = response.data
          this.renderChart()
          this.drawLine()
        })
        .catch(error => {
          console.error('Failed to fetch chart data:', error)
        })
    },
    renderChart() {
      const chart = this.$echarts.init(document.getElementById('myChart'))

      // 配置饼状图数据和样式
      const options = {
        title: {
          text: '缺陷分布饼状图',
          left: 'center'
        },
        tooltip: {
          trigger: 'item'
        },
        legend: {
          top: 0,
          left: 20,
          orient: 'vertical',
          align: 'left'
        },
        series: [
          {
            type: 'pie',
            radius: ['40%', '70%'],
            avoidLabelOverlap: false,
            itemStyle: {
              borderRadius: 10,
              borderColor: '#fff',
              borderWidth: 2
            },
            label: {
              show: false,
              position: 'center'
            },
            emphasis: {
              label: {
                show: true,
                fontSize: 40,
                fontWeight: 'bold'
              }
            },
            labelLine: {
              show: false
            },
            data: this.chartData.map(item => ({
              value: item.value,
              name: item.name
            }))
          }
        ]
      }

      chart.setOption(options)
    },
    drawLine() {
      const lineChart = this.$echarts.init(document.getElementById('lineChart'))

      const option = {
        title: {
          text: '缺陷数量柱状图',
          left: 'center'
        },
        xAxis: {
          type: 'category',
          data: this.chartData.map(item => item.name)
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            name: 'Defects',
            type: 'bar',
            data: this.chartData.map(item => item.value)
          }
        ]
      }

      lineChart.setOption(option)
    }
  }
}
</script>

<style>
.image-container {
  display: flex;
}

#lineChart,
#myChart {
  flex: 1;
  margin-right: 10px;
}
</style>
