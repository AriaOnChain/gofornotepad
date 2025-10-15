// templates.go
package templates

const IndexTemplate = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        * { box-sizing: border-box; margin: 0; padding: 0; }
        body { 
            font-family: 'Microsoft YaHei', sans-serif; 
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 20px;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
            background: white;
            border-radius: 15px;
            box-shadow: 0 20px 40px rgba(0,0,0,0.1);
            overflow: hidden;
        }
        .header {
            background: #4f46e5;
            color: white;
            padding: 30px;
            text-align: center;
        }
        .header h1 {
            margin-bottom: 10px;
        }
        .nav-buttons {
            display: flex;
            gap: 10px;
            margin-top: 15px;
            justify-content: center;
        }
        .nav-btn {
            background: rgba(255,255,255,0.2);
            color: white;
            padding: 8px 16px;
            border: 1px solid rgba(255,255,255,0.3);
            border-radius: 5px;
            text-decoration: none;
            font-size: 12px;
            transition: all 0.3s;
        }
        .nav-btn:hover {
            background: rgba(255,255,255,0.3);
        }
        .stats {
            font-size: 14px;
            opacity: 0.9;
        }
        .content {
            padding: 30px;
        }
        .section {
            margin-bottom: 30px;
            padding: 20px;
            border: 1px solid #e5e7eb;
            border-radius: 10px;
        }
        .form-group {
            margin-bottom: 15px;
        }
        label {
            display: block;
            margin-bottom: 5px;
            font-weight: bold;
            color: #374151;
        }
        input, textarea, select {
            width: 100%;
            padding: 10px;
            border: 1px solid #d1d5db;
            border-radius: 5px;
            font-size: 14px;
            font-family: inherit;
        }
        textarea {
            height: 100px;
            resize: vertical;
        }
        button {
            background: #4f46e5;
            color: white;
            padding: 12px 24px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            font-size: 14px;
            font-weight: bold;
        }
        button:hover {
            background: #4338ca;
        }
        .btn-edit {
            background: #10b981;
        }
        .btn-edit:hover {
            background: #059669;
        }
        .btn-danger {
            background: #dc2626;
        }
        .btn-danger:hover {
            background: #b91c1c;
        }
        .record-item {
            padding: 15px;
            border: 1px solid #e5e7eb;
            border-radius: 8px;
            margin-bottom: 15px;
            background: #f9fafb;
        }
        .record-header {
            display: flex;
            justify-content: space-between;
            align-items: flex-start;
            margin-bottom: 10px;
        }
        .record-title {
            font-weight: bold;
            font-size: 16px;
            color: #1f2937;
        }
        .record-content {
            color: #6b7280;
            line-height: 1.5;
            margin-bottom: 10px;
            white-space: pre-wrap;
        }
        .record-meta {
            font-size: 12px;
            color: #9ca3af;
        }
        .record-actions {
            margin-top: 10px;
            display: flex;
            gap: 10px;
        }
        .search-box {
            display: flex;
            gap: 10px;
            margin-bottom: 20px;
        }
        .search-box input {
            flex: 1;
        }
        .alert {
            padding: 12px;
            border-radius: 5px;
            margin-bottom: 20px;
        }
        .alert-success {
            background: #d1fae5;
            color: #065f46;
            border: 1px solid #a7f3d0;
        }
        .alert-error {
            background: #fee2e2;
            color: #991b1b;
            border: 1px solid #fecaca;
        }
        .alert-info {
            background: #dbeafe;
            color: #1e40af;
            border: 1px solid #93c5fd;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>{{.Title1}}</h1>
            <div class="stats">共 {{.TotalCount}} 条记录</div>
            <div class="nav-buttons">
                <a href="/stats" class="nav-btn">📊 统计</a>
                <a href="/about" class="nav-btn">ℹ️ 关于</a>
            </div>
        </div>
        
        <div class="content">
            <!-- 消息提示 -->
            {{if .SearchQuery}}
            <div class="alert alert-info">
                搜索关键词: "{{.SearchQuery}}"
                <a href="/" style="float:right; color: inherit;">显示全部</a>
            </div>
            {{end}}
            
            {{if .Success}}
            <div class="alert alert-success">{{.Success}}</div>
            {{end}}
            
            {{if .Error}}
            <div class="alert alert-error">{{.Error}}</div>
            {{end}}
            
            <!-- 搜索框 -->
            <form action="/search" method="GET" class="search-box">
                <input type="text" name="q" placeholder="搜索记录..." value="{{.SearchQuery}}">
                <button type="submit">搜索</button>
            </form>
            
            <!-- 添加记录表单 -->
            <div class="section">
                <h2>添加新记录</h2>
                <form action="/add" method="POST">
                    <div class="form-group">
                        <label for="title">标题 *</label>
                        <input type="text" id="title" name="title" required>
                    </div>
                    <div class="form-group">
                        <label for="content">内容</label>
                        <textarea id="content" name="content" placeholder="记录详细信息..."></textarea>
                    </div>
                    <button type="submit">保存记录</button>
                </form>
            </div>
            
            <!-- 记录列表 -->
            <div class="section">
                <h2>记录列表</h2>
                {{if .Records}}
                    {{range .Records}}
                    <div class="record-item">
                        <div class="record-header">
                            <div class="record-title">{{.Title}}</div>
                        </div>
                        {{if .Content}}
                        <div class="record-content">{{.Content}}</div>
                        {{end}}
                        <div class="record-meta">
                            创建时间: {{.CreatedAt.Format "2006-01-02 15:04"}}
                            {{if .UpdatedAt}}
                            | 最后更新: {{.UpdatedAt.Format "2006-01-02 15:04"}}
                            {{end}}
                        </div>
                        <div class="record-actions">
                            <a href="/edit/{{.ID}}">
                                <button type="button" class="btn-edit">编辑</button>
                            </a>
                            <form action="/delete" method="POST" style="display: inline;">
                                <input type="hidden" name="id" value="{{.ID}}">
                                <button type="submit" class="btn-danger" onclick="return confirm('确定删除这条记录吗？')">删除</button>
                            </form>
                        </div>
                    </div>
                    {{end}}
                {{else}}
                    <p style="text-align: center; color: #6b7280; padding: 40px;">
                        {{if .SearchQuery}}
                        没有找到匹配的记录
                        {{else}}
                        还没有任何记录，开始添加第一条吧！
                        {{end}}
                    </p>
                {{end}}
            </div>
        </div>
    </div>
</body>
</html>`

const AboutTemplate = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        * { box-sizing: border-box; margin: 0; padding: 0; }
        body { 
            font-family: 'Microsoft YaHei', sans-serif; 
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 20px;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
            background: white;
            border-radius: 15px;
            box-shadow: 0 20px 40px rgba(0,0,0,0.1);
            overflow: hidden;
        }
        .header {
            background: #10b981;
            color: white;
            padding: 30px;
            text-align: center;
        }
        .content {
            padding: 40px;
        }
        .feature-list {
            list-style: none;
            margin: 20px 0;
        }
        .feature-list li {
            padding: 10px 0;
            border-bottom: 1px solid #e5e7eb;
        }
        .feature-list li:before {
            content: "✓ ";
            color: #10b981;
            font-weight: bold;
        }
        .stats-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
            gap: 20px;
            margin: 30px 0;
        }
        .stat-card {
            background: #f8fafc;
            padding: 20px;
            border-radius: 10px;
            text-align: center;
            border-left: 4px solid #10b981;
        }
        .stat-number {
            font-size: 2em;
            font-weight: bold;
            color: #1f2937;
        }
        .stat-label {
            color: #6b7280;
            font-size: 0.9em;
        }
        .nav-buttons {
            display: flex;
            gap: 10px;
            margin-top: 30px;
        }
        .btn {
            padding: 12px 24px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            text-decoration: none;
            display: inline-block;
            text-align: center;
        }
        .btn-primary {
            background: #4f46e5;
            color: white;
        }
        .btn-secondary {
            background: #6b7280;
            color: white;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>关于 {{.AppName}}</h1>
            <p>版本 {{.Version}} - 记录你的每一个想法</p>
        </div>
        
        <div class="content">
            <h2>应用介绍</h2>
            <p>这是一个简单易用的记录本应用，帮助你记录零碎的想法、待办事项和重要信息。</p>
            
            <h3>主要功能</h3>
            <ul class="feature-list">
                <li>添加和编辑记录</li>
                <li>删除不需要的记录</li>
                <li>搜索特定记录</li>
                <li>查看数据统计</li>
                <li>响应式设计</li>
            </ul>
            
            <div class="stats-grid">
                <div class="stat-card">
                    <div class="stat-number">{{.TotalRecords}}</div>
                    <div class="stat-label">总记录数</div>
                </div>
                <div class="stat-card">
                    <div class="stat-number">5</div>
                    <div class="stat-label">功能模块</div>
                </div>
                <div class="stat-card">
                    <div class="stat-number">1.1.0</div>
                    <div class="stat-label">当前版本</div>
                </div>
            </div>
            
            <div class="nav-buttons">
                <a href="/" class="btn btn-primary">返回首页</a>
                <a href="/stats" class="btn btn-secondary">查看统计</a>
            </div>
        </div>
    </div>
</body>
</html>`

const StatsTemplate = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        * { box-sizing: border-box; margin: 0; padding: 0; }
        body { 
            font-family: 'Microsoft YaHei', sans-serif; 
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 20px;
        }
        .container {
            max-width: 1000px;
            margin: 0 auto;
            background: white;
            border-radius: 15px;
            box-shadow: 0 20px 40px rgba(0,0,0,0.1);
            overflow: hidden;
        }
        .header {
            background: #f59e0b;
            color: white;
            padding: 30px;
            text-align: center;
        }
        .content {
            padding: 40px;
        }
        .stats-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 20px;
            margin: 30px 0;
        }
        .stat-card {
            background: #f8fafc;
            padding: 25px;
            border-radius: 10px;
            text-align: center;
            border-left: 4px solid #f59e0b;
        }
        .stat-number {
            font-size: 2.5em;
            font-weight: bold;
            color: #1f2937;
        }
        .stat-label {
            color: #6b7280;
            font-size: 0.9em;
            margin-top: 5px;
        }
        .record-stats {
            margin: 30px 0;
        }
        .stat-item {
            display: flex;
            justify-content: space-between;
            padding: 15px;
            border-bottom: 1px solid #e5e7eb;
        }
        .nav-buttons {
            display: flex;
            gap: 10px;
            margin-top: 30px;
        }
        .btn {
            padding: 12px 24px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            text-decoration: none;
            display: inline-block;
            text-align: center;
        }
        .btn-primary {
            background: #4f46e5;
            color: white;
        }
        .btn-secondary {
            background: #6b7280;
            color: white;
        }
        .chart {
            background: #f8fafc;
            padding: 20px;
            border-radius: 10px;
            margin: 20px 0;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>数据统计</h1>
            <p>全面了解你的记录情况</p>
        </div>
        
        <div class="content">
            <h2>概览统计</h2>
            <div class="stats-grid">
                <div class="stat-card">
                    <div class="stat-number">{{.Stats.TotalRecords}}</div>
                    <div class="stat-label">总记录数</div>
                </div>
                <div class="stat-card">
                    <div class="stat-number">{{.Stats.TodayRecords}}</div>
                    <div class="stat-label">今日新增</div>
                </div>
                <div class="stat-card">
                    <div class="stat-number">{{.Stats.WeekRecords}}</div>
                    <div class="stat-label">本周新增</div>
                </div>
                <div class="stat-card">
                    <div class="stat-number">{{.Stats.MonthRecords}}</div>
                    <div class="stat-label">本月新增</div>
                </div>
            </div>
            
            <h2>内容统计</h2>
            <div class="stats-grid">
                <div class="stat-card">
                    <div class="stat-number">{{.Stats.TotalChars}}</div>
                    <div class="stat-label">总字符数</div>
                </div>
                <div class="stat-card">
                    <div class="stat-number">{{.Stats.AverageChars}}</div>
                    <div class="stat-label">平均字符数</div>
                </div>
            </div>
            
            {{if .Stats.LongestRecord}}
            <div class="record-stats">
                <h3>记录详情</h3>
                <div class="stat-item">
                    <span><strong>最长的记录:</strong> {{.Stats.LongestRecord.Title}}</span>
                    <span>{{len .Stats.LongestRecord.Content}} 字符</span>
                </div>
                {{if .Stats.ShortestRecord.Content}}
                <div class="stat-item">
                    <span><strong>最短的记录:</strong> {{.Stats.ShortestRecord.Title}}</span>
                    <span>{{len .Stats.ShortestRecord.Content}} 字符</span>
                </div>
                {{end}}
            </div>
            {{end}}
            
            <div class="nav-buttons">
                <a href="/" class="btn btn-primary">返回首页</a>
                <a href="/about" class="btn btn-secondary">关于我们</a>
            </div>
        </div>
    </div>
</body>
</html>`

const LinksTemplate = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        * { box-sizing: border-box; margin: 0; padding: 0; }
        body { 
            font-family: 'Microsoft YaHei', sans-serif; 
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 20px;
        }
        .container {
            max-width: 900px;
            margin: 0 auto;
            background: white;
            border-radius: 15px;
            box-shadow: 0 20px 40px rgba(0,0,0,0.1);
            overflow: hidden;
        }
        .header {
            background: #8b5cf6;
            color: white;
            padding: 30px;
            text-align: center;
        }
        .header h1 {
            margin-bottom: 10px;
        }
        .stats {
            font-size: 14px;
            opacity: 0.9;
        }
        .nav-buttons {
            display: flex;
            gap: 10px;
            margin-top: 15px;
            justify-content: center;
        }
        .nav-btn {
            background: rgba(255,255,255,0.2);
            color: white;
            padding: 8px 16px;
            border: 1px solid rgba(255,255,255,0.3);
            border-radius: 5px;
            text-decoration: none;
            font-size: 12px;
            transition: all 0.3s;
        }
        .nav-btn:hover {
            background: rgba(255,255,255,0.3);
        }
        .content {
            padding: 30px;
        }
        .section {
            margin-bottom: 30px;
            padding: 20px;
            border: 1px solid #e5e7eb;
            border-radius: 10px;
        }
        .form-group {
            margin-bottom: 15px;
        }
        label {
            display: block;
            margin-bottom: 5px;
            font-weight: bold;
            color: #374151;
        }
        input, textarea {
            width: 100%;
            padding: 10px;
            border: 1px solid #d1d5db;
            border-radius: 5px;
            font-size: 14px;
            font-family: inherit;
        }
        textarea {
            height: 80px;
            resize: vertical;
        }
        button {
            background: #8b5cf6;
            color: white;
            padding: 12px 24px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            font-size: 14px;
            font-weight: bold;
            transition: all 0.3s;
        }
        button:hover {
            background: #7c3aed;
        }
        .btn-edit {
            background: #10b981;
        }
        .btn-edit:hover {
            background: #059669;
        }
        .btn-danger {
            background: #dc2626;
        }
        .btn-danger:hover {
            background: #b91c1c;
        }
        .btn-link {
            background: #3b82f6;
        }
        .btn-link:hover {
            background: #2563eb;
        }
        .record-item {
            padding: 15px;
            border: 1px solid #e5e7eb;
            border-radius: 8px;
            margin-bottom: 15px;
            background: #f9fafb;
            transition: all 0.3s;
        }
        .record-item:hover {
            box-shadow: 0 2px 8px rgba(0,0,0,0.1);
        }
        .record-header {
            display: flex;
            justify-content: space-between;
            align-items: flex-start;
            margin-bottom: 10px;
        }
        .record-title {
            font-weight: bold;
            font-size: 16px;
            color: #1f2937;
        }
        .record-content {
            color: #6b7280;
            line-height: 1.5;
            margin-bottom: 10px;
            white-space: pre-wrap;
        }
        .record-link {
            margin: 10px 0;
        }
        .record-link a {
            color: #3b82f6;
            text-decoration: none;
            word-break: break-all;
        }
        .record-link a:hover {
            text-decoration: underline;
        }
        .record-meta {
            font-size: 12px;
            color: #9ca3af;
        }
        .record-actions {
            margin-top: 10px;
            display: flex;
            gap: 10px;
            flex-wrap: wrap;
        }
        .search-box {
            display: flex;
            gap: 10px;
            margin-bottom: 20px;
        }
        .search-box input {
            flex: 1;
        }
        .alert {
            padding: 12px;
            border-radius: 5px;
            margin-bottom: 20px;
        }
        .alert-success {
            background: #d1fae5;
            color: #065f46;
            border: 1px solid #a7f3d0;
        }
        .alert-error {
            background: #fee2e2;
            color: #991b1b;
            border: 1px solid #fecaca;
        }
        .alert-info {
            background: #dbeafe;
            color: #1e40af;
            border: 1px solid #93c5fd;
        }

        /* 模态框样式 */
        .modal {
            display: none;
            position: fixed;
            z-index: 1000;
            left: 0;
            top: 0;
            width: 100%;
            height: 100%;
            background-color: rgba(0,0,0,0.5);
        }
        .modal-content {
            background-color: white;
            margin: 5% auto;
            padding: 0;
            border-radius: 15px;
            width: 90%;
            max-width: 600px;
            box-shadow: 0 20px 40px rgba(0,0,0,0.3);
            animation: modalSlideIn 0.3s ease-out;
        }
        @keyframes modalSlideIn {
            from { transform: translateY(-50px); opacity: 0; }
            to { transform: translateY(0); opacity: 1; }
        }
        .modal-header {
            background: #8b5cf6;
            color: white;
            padding: 20px 30px;
            border-radius: 15px 15px 0 0;
        }
        .modal-body {
            padding: 30px;
        }
        .modal-footer {
            padding: 20px 30px;
            border-top: 1px solid #e5e7eb;
            display: flex;
            gap: 10px;
            justify-content: flex-end;
        }
        .close {
            color: white;
            float: right;
            font-size: 28px;
            font-weight: bold;
            cursor: pointer;
            line-height: 1;
        }
        .close:hover {
            color: #ddd6fe;
        }
        .loading {
            display: none;
            text-align: center;
            padding: 20px;
            color: #6b7280;
        }
        .url-preview {
            background: #f8fafc;
            padding: 8px 12px;
            border-radius: 4px;
            font-size: 12px;
            color: #6b7280;
            margin-top: 5px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>{{.Title1}}</h1>
            <div class="stats">共 {{.TotalCount}} 个链接</div>
            <div class="nav-buttons">
                <a href="/" class="nav-btn">📝 想法记录</a>
                <a href="/stats" class="nav-btn">📊 统计</a>
                <a href="/about" class="nav-btn">ℹ️ 关于</a>
            </div>
        </div>
        
        <div class="content">
            <!-- 消息提示 -->
            {{if .SearchQuery}}
            <div class="alert alert-info">
                搜索关键词: "{{.SearchQuery}}"
                <a href="/links" style="float:right; color: inherit;">显示全部</a>
            </div>
            {{end}}
            
            {{if .Success}}
            <div class="alert alert-success">{{.Success}}</div>
            {{end}}
            
            {{if .Error}}
            <div class="alert alert-error">{{.Error}}</div>
            {{end}}
            
            <!-- 搜索框 -->
            <form action="/links" method="GET" class="search-box">
                <input type="text" name="q" placeholder="搜索链接、标题或内容..." value="{{.SearchQuery}}">
                <button type="submit">搜索</button>
            </form>
            
            <!-- 添加链接表单 -->
            <div class="section">
                <h2>添加新链接</h2>
                <form action="/links/add" method="POST">
                    <div class="form-group">
                        <label for="title">标题</label>
                        <input type="text" id="title" name="title" placeholder="链接标题...">
                    </div>
                    <div class="form-group">
                        <label for="link">链接 *</label>
                        <input type="url" id="link" name="link" placeholder="https://example.com" required>
                        <div class="url-preview">支持 http:// 和 https:// 开头的链接</div>
                    </div>
                    <div class="form-group">
                        <label for="content">描述/备注</label>
                        <textarea id="content" name="content" placeholder="链接描述或备注..."></textarea>
                    </div>
                    <button type="submit">保存链接</button>
                </form>
            </div>
            
            <!-- 链接列表 -->
            <div class="section">
                <h2>链接列表</h2>
                <div id="records-container">
                    {{if .Records}}
                        {{range .Records}}
                        <div class="record-item" id="record-{{.ID}}">
                            <div class="record-header">
                                <div class="record-title">{{.Title}}</div>
                            </div>
                            
                            {{if .Link}}
                            <div class="record-link">
                                <a href="{{.Link}}" target="_blank" rel="noopener noreferrer">
                                    🔗 {{.Link}}
                                </a>
                            </div>
                            {{end}}
                            
                            {{if .Content}}
                            <div class="record-content">{{.Content}}</div>
                            {{end}}
                            
                            <div class="record-meta">
                                创建时间: {{.CreatedAt.Format "2006-01-02 15:04"}}
                                {{if .UpdatedAt}}
                                | 最后更新: {{.UpdatedAt.Format "2006-01-02 15:04"}}
                                {{end}}
                            </div>
                            
                            <div class="record-actions">
                                {{if .Link}}
                                <a href="{{.Link}}" target="_blank" rel="noopener noreferrer">
                                    <button type="button" class="btn-link">访问链接</button>
                                </a>
                                {{end}}
                                <button type="button" class="btn-edit" onclick="openEditModal({{.ID}}, '{{.Title}}', '{{.Content}}', '{{.Link}}')">编辑</button>
                                <form action="/links/delete" method="POST" style="display: inline;">
                                    <input type="hidden" name="id" value="{{.ID}}">
                                    <button type="submit" class="btn-danger" onclick="return confirm('确定删除这个链接吗？')">删除</button>
                                </form>
                            </div>
                        </div>
                        {{end}}
                    {{else}}
                        <p style="text-align: center; color: #6b7280; padding: 40px;">
                            {{if .SearchQuery}}
                            没有找到匹配的链接
                            {{else}}
                            还没有任何链接收藏，开始添加第一个吧！
                            {{end}}
                        </p>
                    {{end}}
                </div>
            </div>
        </div>
    </div>

    <!-- 编辑模态框 -->
    <div id="editModal" class="modal">
        <div class="modal-content">
            <div class="modal-header">
                <h2 style="margin: 0;">编辑链接</h2>
                <span class="close" onclick="closeEditModal()">&times;</span>
            </div>
            <div class="modal-body">
                <form id="editForm">
                    <input type="hidden" id="editRecordId">
                    <div class="form-group">
                        <label for="editTitle">标题 *</label>
                        <input type="text" id="editTitle" name="title" required>
                    </div>
                    <div class="form-group">
                        <label for="editLink">链接 *</label>
                        <input type="url" id="editLink" name="link" required>
                    </div>
                    <div class="form-group">
                        <label for="editContent">描述/备注</label>
                        <textarea id="editContent" name="content" placeholder="链接描述或备注..." rows="4"></textarea>
                    </div>
                </form>
                <div class="loading" id="editLoading">
                    更新中...
                </div>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn-secondary" onclick="closeEditModal()">取消</button>
                <button type="button" class="btn-primary" onclick="updateRecord()">更新链接</button>
            </div>
        </div>
    </div>

    <script>
        // 打开编辑模态框
        function openEditModal(id, title, content, link) {
            document.getElementById('editRecordId').value = id;
            document.getElementById('editTitle').value = title;
            document.getElementById('editContent').value = content;
            document.getElementById('editLink').value = link;
            document.getElementById('editModal').style.display = 'block';
        }

        // 关闭编辑模态框
        function closeEditModal() {
            document.getElementById('editModal').style.display = 'none';
            document.getElementById('editLoading').style.display = 'none';
        }

        // 点击模态框外部关闭
        window.onclick = function(event) {
            const modal = document.getElementById('editModal');
            if (event.target === modal) {
                closeEditModal();
            }
        }

        // 更新记录
        async function updateRecord() {
            const id = document.getElementById('editRecordId').value;
            const title = document.getElementById('editTitle').value.trim();
            const content = document.getElementById('editContent').value.trim();
            const link = document.getElementById('editLink').value.trim();

            if (!title) {
                alert('标题不能为空');
                return;
            }

            if (!link) {
                alert('链接不能为空');
                return;
            }

            // 显示加载中
            document.getElementById('editLoading').style.display = 'block';

            try {
                const response = await fetch('/links/api/update', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        id: parseInt(id),
                        title: title,
                        content: content,
                        link: link
                    })
                });

                const result = await response.json();

                if (result.success) {
                    // 重新加载页面以获取最新数据
                    location.reload();
                } else {
                    throw new Error(result.message || '更新失败');
                }
            } catch (error) {
                console.error('更新失败:', error);
                alert('更新失败: ' + error.message);
            } finally {
                document.getElementById('editLoading').style.display = 'none';
            }
        }
    </script>
</body>
</html>`
