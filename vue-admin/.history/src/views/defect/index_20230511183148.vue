<template>
  <div class='app-container'>
    <el-input
      v-model='filterText'
      placeholder='Filter keyword'
      style='margin-bottom: 30px'
    />

    <el-tree
      ref='tree2'
      :data='data2'
      :props='defaultProps'
      :filter-node-method='filterNode'
      class='filter-tree'
      default-expand-all
    >
      <template v-slot='{ node: { data } }'>
        <div class='demo-image__placeholder'>
          <div class='block'>
            <span class='demonstration'>原图</span>
            <el-image :src='data.src1'></el-image>
          </div>
          <div class='block'>
            <span class='demonstration'>检测图</span>
            <el-image :src='data.src2'>
              <div slot='placeholder' class='image-slot'>
                加载中<span class='dot'>...</span>
              </div>
            </el-image>
          </div>
        </div>
      </template>
    </el-tree>
  </div>
</template>

<script>
export default {
  data() {
    return {
      filterText: '',
      data2: [
        {
          id: 1,
          label: 'Level one 1',
          src1: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139735_orignal.jpg',
          src2: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139737_result.jpg',
          children: [
            {
              id: 4,
              label: 'Level two 1-1',
              src1: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139735_orignal.jpg',
              src2: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139737_result.jpg',
              children: [
                {
                  id: 9,
                  label: 'Level three 1-1-1',
                  src1: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139735_orignal.jpg',
                  src2: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139737_result.jpg',
                },
                {
                  id: 10,
                  label: 'Level three 1-1-2',
                  src1: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139735_orignal.jpg',
                  src2: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139737_result.jpg',
                },
              ],
            },
          ],
        },
        {
          id: 2,
          label: 'Level one 2',
          src1: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139735_orignal.jpg',
          src2: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139737_result.jpg',
          children: [
            {
              id: 5,
              label: 'Level two 2-1',
              src1: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139735_orignal.jpg',
              src2: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139737_result.jpg',
            },
            {
              id: 6,
              label: 'Level two 2-2',
              src1: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139735_orignal.jpg',
              src2: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139737_result.jpg',
            },
          ],
        },
        {
          id: 3,
          label: 'Level one 3',
          src1: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139735_orignal.jpg',
          src2: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139737_result.jpg',
          children: [
            {
              id: 7,
              label: 'Level two 3-1',
              src1: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139735_orignal.jpg',
              src2: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139737_result.jpg',
            },
            {
              id: 8,
              label: 'Level two 3-2',
              src1: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139735_orignal.jpg',
              src2: 'https://echo-img-1305917469.cos.ap-nanjing.myqcloud.com/defect/1683139737_result.jpg',
            },
          ],
        },
      ],
      defaultProps: {
        children: 'children',
        label: 'label',
      },
    };
  },
  watch: {
    filterText(val) {
      this.$refs.tree2.filter(val);
    },
  },
  methods: {
    filterNode(value, data) {
      if (!value) return true;
      return data.label.indexOf(value) !== -1;
    },
  },
};
</script>

<style>
.demo-image__placeholder {
  display: flex;
  align-items: center;
}

.block {
  margin-right: 20px;
}

.demonstration {
  font-size: 14px;
  font-weight: bold;
  margin-right: 10px;
}

.image-slot {
  display: inline-block;
  width: 70px;
  text-align: center;
  font-size: 12px;
  color: #999;
  border: 1px solid #ccc;
  border-radius: 4px;
  padding: 5px;
  margin-right: 10px;
}
</style>
