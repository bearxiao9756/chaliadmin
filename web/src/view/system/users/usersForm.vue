
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户ID:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入用户ID" />
</el-form-item>
        <el-form-item label="用户类型:" prop="userType">
    <el-input v-model.number="formData.userType" :clearable="true" placeholder="请输入用户类型" />
</el-form-item>
        <el-form-item label="accessHash:" prop="accessHash">
    <el-input v-model.number="formData.accessHash" :clearable="true" placeholder="请输入accessHash" />
</el-form-item>
        <el-form-item label="secretKey:" prop="secretKeyId">
    <el-input v-model.number="formData.secretKeyId" :clearable="true" placeholder="请输入secretKey" />
</el-form-item>
        <el-form-item label="姓名-姓名:" prop="firstName">
    <el-input v-model="formData.firstName" :clearable="true" placeholder="请输入姓名-姓名" />
</el-form-item>
        <el-form-item label="姓名-姓:" prop="lastName">
    <el-input v-model="formData.lastName" :clearable="true" placeholder="请输入姓名-姓" />
</el-form-item>
        <el-form-item label="姓名-氏:" prop="username">
    <el-input v-model="formData.username" :clearable="true" placeholder="请输入姓名-氏" />
</el-form-item>
        <el-form-item label="手机:" prop="phone">
    <el-input v-model="formData.phone" :clearable="true" placeholder="请输入手机" />
</el-form-item>
        <el-form-item label="国家编码:" prop="countryCode">
    <el-input v-model="formData.countryCode" :clearable="true" placeholder="请输入国家编码" />
</el-form-item>
        <el-form-item label="是否验证:" prop="verified">
    <el-switch v-model="formData.verified" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="是否支持:" prop="support">
    <el-switch v-model="formData.support" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="是否欺诈:" prop="scam">
    <el-switch v-model="formData.scam" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="是否假冒:" prop="fake">
    <el-switch v-model="formData.fake" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="是否特权:" prop="premium">
    <el-switch v-model="formData.premium" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="特权时间:" prop="premiumExpireDate">
    <el-input v-model.number="formData.premiumExpireDate" :clearable="true" placeholder="请输入特权时间" />
</el-form-item>
        <el-form-item label="个人简介:" prop="about">
    <el-input v-model="formData.about" :clearable="true" placeholder="请输入个人简介" />
</el-form-item>
        <el-form-item label="用户状态:" prop="state">
    <el-input v-model.number="formData.state" :clearable="true" placeholder="请输入用户状态" />
</el-form-item>
        <el-form-item label="是否机器人:" prop="isBot">
    <el-switch v-model="formData.isBot" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="自动删除时间:" prop="accountDaysTtl">
    <el-input v-model.number="formData.accountDaysTtl" :clearable="true" placeholder="请输入自动删除时间" />
</el-form-item>
        <el-form-item label="头像ID:" prop="photoId">
    <el-input v-model.number="formData.photoId" :clearable="true" placeholder="请输入头像ID" />
</el-form-item>
        <el-form-item label="是否封禁:" prop="restricted">
    <el-switch v-model="formData.restricted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="封禁原因:" prop="restrictionReason">
    <el-input v-model="formData.restrictionReason" :clearable="true" placeholder="请输入封禁原因" />
</el-form-item>
        <el-form-item label="是否陌生人静音:" prop="archiveAndMuteNewNoncontactPeers">
    <el-switch v-model="formData.archiveAndMuteNewNoncontactPeers" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="表情文档ID:" prop="emojiStatusDocumentId">
    <el-input v-model.number="formData.emojiStatusDocumentId" :clearable="true" placeholder="请输入表情文档ID" />
</el-form-item>
        <el-form-item label="表情到期时间:" prop="emojiStatusUntil">
    <el-input v-model.number="formData.emojiStatusUntil" :clearable="true" placeholder="请输入表情到期时间" />
</el-form-item>
        <el-form-item label="最新的故事 ID:" prop="storiesMaxId">
    <el-input v-model.number="formData.storiesMaxId" :clearable="true" placeholder="请输入最新的故事 ID" />
</el-form-item>
        <el-form-item label="主颜色:" prop="color">
    <el-input v-model.number="formData.color" :clearable="true" placeholder="请输入主颜色" />
</el-form-item>
        <el-form-item label="表情背景色ID:" prop="colorBackgroundEmojiId">
    <el-input v-model.number="formData.colorBackgroundEmojiId" :clearable="true" placeholder="请输入表情背景色ID" />
</el-form-item>
        <el-form-item label="个人页面主色:" prop="profileColor">
    <el-input v-model.number="formData.profileColor" :clearable="true" placeholder="请输入个人页面主色" />
</el-form-item>
        <el-form-item label="个人页面表情背景色ID:" prop="profileColorBackgroundEmojiId">
    <el-input v-model.number="formData.profileColorBackgroundEmojiId" :clearable="true" placeholder="请输入个人页面表情背景色ID" />
</el-form-item>
        <el-form-item label="生日:" prop="birthday">
    <el-input v-model="formData.birthday" :clearable="true" placeholder="请输入生日" />
</el-form-item>
        <el-form-item label="个人频道ID:" prop="personalChannelId">
    <el-input v-model.number="formData.personalChannelId" :clearable="true" placeholder="请输入个人频道ID" />
</el-form-item>
        <el-form-item label="Token有效期:" prop="authorizationTtlDays">
    <el-input v-model.number="formData.authorizationTtlDays" :clearable="true" placeholder="请输入Token有效期" />
</el-form-item>
        <el-form-item label="是否删除:" prop="deleted">
    <el-switch v-model="formData.deleted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="删除原因:" prop="deleteReason">
    <el-input v-model="formData.deleteReason" :clearable="true" placeholder="请输入删除原因" />
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
  createUsers,
  updateUsers,
  findUsers
} from '@/api/system/users'

defineOptions({
    name: 'UsersForm'
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
            userType: undefined,
            accessHash: undefined,
            secretKeyId: undefined,
            firstName: '',
            lastName: '',
            username: '',
            phone: '',
            countryCode: '',
            verified: false,
            support: false,
            scam: false,
            fake: false,
            premium: false,
            premiumExpireDate: undefined,
            about: '',
            state: undefined,
            isBot: false,
            accountDaysTtl: undefined,
            photoId: undefined,
            restricted: false,
            restrictionReason: '',
            archiveAndMuteNewNoncontactPeers: false,
            emojiStatusDocumentId: undefined,
            emojiStatusUntil: undefined,
            storiesMaxId: undefined,
            color: undefined,
            colorBackgroundEmojiId: undefined,
            profileColor: undefined,
            profileColorBackgroundEmojiId: undefined,
            birthday: '',
            personalChannelId: undefined,
            authorizationTtlDays: undefined,
            deleted: false,
            deleteReason: '',
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
      const res = await findUsers({ ID: route.query.id })
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
               res = await createUsers(formData.value)
               break
             case 'update':
               res = await updateUsers(formData.value)
               break
             default:
               res = await createUsers(formData.value)
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
