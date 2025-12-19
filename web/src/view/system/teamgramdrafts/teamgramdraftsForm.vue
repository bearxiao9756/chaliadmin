
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="编号:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入编号" />
</el-form-item>
        <el-form-item label="用户身份:" prop="userId">
    <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入用户身份" />
</el-form-item>
        <el-form-item label="会话ID:" prop="peerDialogId">
    <el-input v-model.number="formData.peerDialogId" :clearable="true" placeholder="请输入会话ID" />
</el-form-item>
        <el-form-item label="草稿类型:" prop="draftType">
    <el-input v-model.number="formData.draftType" :clearable="true" placeholder="请输入草稿类型" />
</el-form-item>
        <el-form-item label="草稿消息数据:" prop="draftMessageData">
    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.draftMessageData 后端会按照json的类型进行存取
    {{ formData.draftMessageData }}
</el-form-item>
        <el-form-item label="date2字段:" prop="date2">
    <el-input v-model.number="formData.date2" :clearable="true" placeholder="请输入date2字段" />
</el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="更新时间:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createTeamgramDrafts,
  updateTeamgramDrafts,
  findTeamgramDrafts
} from '@/api/system/teamgramdrafts'

defineOptions({
    name: 'TeamgramDraftsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            id: undefined,
            userId: undefined,
            peerDialogId: undefined,
            draftType: undefined,
            draftMessageData: {},
            date2: undefined,
            createdAt: new Date(),
            updatedAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTeamgramDrafts({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createTeamgramDrafts(formData.value)
               break
             case 'update':
               res = await updateTeamgramDrafts(formData.value)
               break
             default:
               res = await createTeamgramDrafts(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
