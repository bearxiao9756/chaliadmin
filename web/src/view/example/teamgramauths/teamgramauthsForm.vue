
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="编号:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入编号" />
</el-form-item>
        <el-form-item label="authKeyId:" prop="authKeyId">
    <el-input v-model.number="formData.authKeyId" :clearable="true" placeholder="请输入authKeyId" />
</el-form-item>
        <el-form-item label="设备类型:" prop="langPack">
    <el-input v-model="formData.langPack" :clearable="true" placeholder="请输入设备类型" />
</el-form-item>
        <el-form-item label="apiId:" prop="apiId">
    <el-input v-model.number="formData.apiId" :clearable="true" placeholder="请输入apiId" />
</el-form-item>
        <el-form-item label="设备IP:" prop="clientIp">
    <el-input v-model="formData.clientIp" :clearable="true" placeholder="请输入设备IP" />
</el-form-item>
        <el-form-item label="设备名称:" prop="deviceModel">
    <el-input v-model="formData.deviceModel" :clearable="true" placeholder="请输入设备名称" />
</el-form-item>
        <el-form-item label="系统语言代码:" prop="systemLangCode">
    <el-input v-model="formData.systemLangCode" :clearable="true" placeholder="请输入系统语言代码" />
</el-form-item>
        <el-form-item label="系统版本:" prop="systemVersion">
    <el-input v-model="formData.systemVersion" :clearable="true" placeholder="请输入系统版本" />
</el-form-item>
        <el-form-item label="软件版本:" prop="appVersion">
    <el-input v-model="formData.appVersion" :clearable="true" placeholder="请输入软件版本" />
</el-form-item>
        <el-form-item label="代理:" prop="proxy">
    <el-input v-model="formData.proxy" :clearable="true" placeholder="请输入代理" />
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
  createAuths,
  updateAuths,
  findAuths
} from '@/api/example/teamgramauths'

defineOptions({
    name: 'AuthsForm'
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
            authKeyId: undefined,
            langPack: '',
            apiId: undefined,
            clientIp: '',
            deviceModel: '',
            systemLangCode: '',
            systemVersion: '',
            appVersion: '',
            proxy: '',
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
      const res = await findAuths({ ID: route.query.id })
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
               res = await createAuths(formData.value)
               break
             case 'update':
               res = await updateAuths(formData.value)
               break
             default:
               res = await createAuths(formData.value)
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
