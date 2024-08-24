<template>
    <div class="breadList">
        <a-breadcrumb>
            <a-breadcrumb-item>博客前端</a-breadcrumb-item>
            <a-breadcrumb-item>文章详情</a-breadcrumb-item>
        </a-breadcrumb>
    </div>
    <div class="content">
        <a-spin class="MyLoading" :loading="isLoading" tip="加载中，请稍候...">
            <div>
                <MdPreview :editorId="String(blog.Id)" :modelValue="blog.content" />
            </div>
        </a-spin>
    </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { onMounted } from 'vue';
import { MdPreview } from 'md-editor-v3';
// preview.css相比style.css少了编辑器那部分样式
import 'md-editor-v3/lib/preview.css';
import { DESCRIBE_BLOG } from '../../../api/vblog/vblog';
var isLoading = ref(true)

var blog = reactive({})
var blog_id = location.search.split('=')[1]

var handleLoad = async () => {
    try {
        isLoading.value = true
        blog = await DESCRIBE_BLOG(blog_id)
        console.log(blog)
    } finally {
        isLoading.value = false
    }
}

onMounted(() => {
    handleLoad()
})

</script>

<style>
.arco-spin.MyLoading {
    display: flex;
}

.md-editor-preview-wrapper {
    max-width: 1300px;
    position: relative;
    flex: 1;
    box-sizing: border-box;
    overflow: auto;
    padding: 10px 20px;
}
</style>