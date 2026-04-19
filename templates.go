package main

const sharedCSS = `
*{box-sizing:border-box;margin:0;padding:0}
body{font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,sans-serif;background:#eef0f4;color:#1a1f36;min-height:100vh}

/* ---- Header ---- */
.header{background:#fff;border-bottom:1px solid #d8dce8;padding:0 28px;display:flex;flex-direction:column;align-items:center;gap:0}
.header-top{width:100%;display:flex;align-items:center;gap:16px;padding:14px 0 10px}
.header-logo{width:48px;height:48px;object-fit:contain;flex-shrink:0}
.header-title{flex:1}
.header-title h1{font-size:18px;font-weight:700;color:#1a1f36;line-height:1.2}
.header-title p{font-size:12px;color:#8a93b2;margin-top:2px}
.header-actions{display:flex;gap:8px;flex-shrink:0;align-items:center}

.search-bar{width:100%;padding:0 0 14px;display:flex;justify-content:center}
.search-form{display:flex;gap:8px;width:100%;max-width:700px;align-items:center}
.search-input{flex:1;background:#f4f6fb;border:1px solid #d0d5e8;border-radius:8px;color:#1a1f36;padding:9px 16px;font-size:14px;outline:none;transition:border-color .15s}
.search-input:focus{border-color:#3b5bdb;background:#fff}
.search-input::placeholder{color:#aab0c8}

/* ---- Buttons ---- */
.btn{background:#3b5bdb;color:#fff;border:none;border-radius:8px;padding:8px 18px;font-size:13px;font-weight:500;cursor:pointer;white-space:nowrap;transition:background .15s}
.btn:hover{background:#2f4ac9}
.btn-outline{background:#fff;color:#3b5bdb;border:1.5px solid #3b5bdb;border-radius:8px;padding:7px 16px;font-size:13px;font-weight:500;cursor:pointer;white-space:nowrap}
.btn-outline:hover{background:#eef1fb}
.btn-sm{background:#f4f6fb;color:#5a6380;border:1px solid #d0d5e8;border-radius:6px;padding:3px 10px;font-size:12px;cursor:pointer;transition:background .15s}
.btn-sm:hover{background:#e4e8f4;color:#1a1f36}
.btn-danger-sm{background:#fff5f5;color:#e03131;border:1px solid #ffc9c9;border-radius:6px;padding:3px 10px;font-size:12px;cursor:pointer}
.btn-danger-sm:hover{background:#ffe3e3}
.btn-ghost{background:transparent;border:1px solid #d0d5e8;color:#8a93b2;border-radius:8px;padding:7px 14px;font-size:13px;cursor:pointer}
.btn-ghost:hover{color:#1a1f36;border-color:#aab0c8}
.clear-link{color:#aab0c8;font-size:13px;white-space:nowrap;cursor:pointer;text-decoration:none}
.clear-link:hover{color:#5a6380}

/* ---- Content ---- */
.content{padding:22px 28px;max-width:1400px;margin:0 auto;display:flex;flex-direction:column;gap:14px}

/* ---- Root department card ---- */
.dept-card{background:#fff;border:1px solid #d8dce8;border-radius:10px;overflow:hidden;box-shadow:0 1px 3px rgba(0,0,0,.06)}

.dept-root-header{background:linear-gradient(90deg,#2c3e7a 0%,#3b5bdb 100%);color:#fff;padding:11px 16px;display:flex;align-items:center;gap:8px}
.dept-root-name{font-size:13px;font-weight:700;letter-spacing:.05em;text-transform:uppercase;flex:1}
.dept-root-actions{display:flex;gap:6px;align-items:center}

/* ---- Sub-department ---- */
.subdept-wrap{border-top:1px solid #edf0f9}
.subdept-header{background:#f0f4ff;padding:8px 16px 8px 28px;display:flex;align-items:center;gap:8px;border-bottom:1px solid #e2e8ff}
.subdept-arrow{color:#7a90d4;font-size:13px}
.subdept-name{font-size:12px;font-weight:600;color:#3b5bdb;letter-spacing:.04em;text-transform:uppercase;flex:1}
.subdept-actions{display:flex;gap:6px;align-items:center}

/* ---- Table ---- */
.table-wrap{overflow-x:auto}
table{width:100%;border-collapse:collapse;min-width:700px}
thead th{background:#f8f9fd;color:#8a93b2;font-size:11px;text-transform:uppercase;letter-spacing:.07em;padding:8px 12px;text-align:left;border-bottom:1px solid #e8eaf4;white-space:nowrap;font-weight:600}
tbody tr:hover{background:#f8f9fd}
tbody td{padding:8px 12px;font-size:13px;border-bottom:1px solid #f0f2f9;vertical-align:middle}
tbody tr:last-child td{border-bottom:none}
.col-room{color:#aab0c8;text-align:center;width:52px;font-size:12px;font-weight:600}
.col-pos{color:#5a6380;max-width:280px;font-size:12px}
.col-name{font-weight:600;color:#1a1f36;white-space:nowrap}
.col-phone{font-family:monospace;color:#2b6cb0;white-space:nowrap;font-size:13px}
.col-email{color:#2f9e88;font-size:12px}
.col-act{width:90px;white-space:nowrap}
.act-row{display:flex;gap:4px;align-items:center}

.empty{color:#c0c6d9;font-size:13px;padding:18px 16px;text-align:center;font-style:italic}
.no-data{color:#c0c6d9;font-size:16px;padding:80px;text-align:center}

/* ---- Footer ---- */
.footer{text-align:center;padding:24px 16px;color:#b0b8d0;font-size:12px;border-top:1px solid #e4e8f4;margin-top:8px}

/* ---- Org filter ---- */
.org-filter{width:100%;padding:0 0 10px;display:flex;justify-content:center}
.org-selector{position:relative;width:100%;max-width:700px;display:flex;align-items:center;gap:8px}
.org-input{flex:1;background:#eef1fb;border:1.5px solid #b8c4f0;border-radius:8px;color:#1a1f36;padding:9px 16px;font-size:14px;font-weight:500;outline:none;transition:border-color .15s,background .15s;cursor:pointer}
.org-input:focus{border-color:#3b5bdb;background:#fff}
.org-input::placeholder{color:#aab0c8;font-weight:400}
.org-dropdown{position:absolute;top:calc(100% + 4px);left:0;right:0;background:#fff;border:1px solid #d0d5e8;border-radius:10px;box-shadow:0 6px 24px rgba(0,0,0,.13);z-index:200;max-height:300px;overflow-y:auto}
.org-item{padding:10px 16px;font-size:14px;cursor:pointer;display:flex;align-items:center;gap:6px;transition:background .1s}
.org-item:hover,.org-item.highlighted{background:#f0f4ff;color:#3b5bdb}
.org-item.selected{font-weight:700;color:#3b5bdb}
.org-item-badge{font-size:10px;color:#aab0c8;margin-left:auto;font-weight:400;background:#f4f6fb;padding:1px 6px;border-radius:4px}
.org-all-item{padding:10px 16px;font-size:13px;cursor:pointer;color:#8a93b2;border-bottom:1px solid #f0f2f9;transition:background .1s}
.org-all-item:hover{background:#f0f4ff}

/* ---- Admin badge ---- */
.admin-badge{background:#e03131;color:#fff;font-size:10px;font-weight:700;letter-spacing:.08em;padding:2px 7px;border-radius:4px}

/* ---- Wide dialogs ---- */
#dlg-orgs{width:720px}

/* ---- Dialog ---- */
dialog{background:#fff;border:1px solid #d0d5e8;border-radius:14px;padding:0;color:#1a1f36;width:540px;max-width:96vw;box-shadow:0 8px 32px rgba(0,0,0,.18)}
dialog::backdrop{background:rgba(30,35,60,.55)}
.dlg-head{padding:20px 24px 14px;border-bottom:1px solid #edf0f9}
.dlg-head h3{font-size:16px;font-weight:700;color:#1a1f36}
.dlg-body{padding:18px 24px;display:flex;flex-direction:column;gap:13px}
.dlg-foot{padding:14px 24px;border-top:1px solid #edf0f9;display:flex;justify-content:flex-end;gap:8px}
label{font-size:12px;color:#5a6380;font-weight:500;display:flex;flex-direction:column;gap:5px}
input[type=text],input[type=email],input[type=number],select{background:#f4f6fb;border:1px solid #d0d5e8;border-radius:7px;color:#1a1f36;padding:8px 11px;font-size:13px;width:100%;outline:none;transition:border-color .15s}
input:focus,select:focus{border-color:#3b5bdb;background:#fff}
.row2{display:grid;grid-template-columns:1fr 1fr;gap:10px}
.row3{display:grid;grid-template-columns:1fr 1fr 1fr;gap:10px}
`

// --- View page ---

const viewPage = `<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8"/>
<meta name="viewport" content="width=device-width,initial-scale=1"/>
<title>Телефонный справочник</title>
<style>` + sharedCSS + `</style>
</head>
<body>
<div class="header">
  <div class="header-top">
    <img class="header-logo" src="/logo" onerror="this.style.display='none'" alt=""/>
    <div class="header-title">
      <h1>{{if .SiteTitle}}{{.SiteTitle}}{{else}}Телефонный справочник{{end}}</h1>
      {{range .HeaderLines}}<p>{{.}}</p>{{end}}
    </div>
    <div class="header-actions">
      <a href="/admin"><button class="btn-ghost" type="button">Администрирование</button></a>
    </div>
  </div>
  {{if .Orgs}}
  <div class="org-filter">
    <div class="org-selector">
      <input id="org-search" class="org-input" type="text"
             value="{{.OrgName}}"
             placeholder="Фильтр по организации..."
             autocomplete="off" readonly/>
      <div id="org-dropdown" class="org-dropdown" style="display:none"></div>
      {{if .CurrentOrg}}<a class="clear-link" href="/?{{if .Search}}q={{.Search}}{{end}}">&#10005; Все</a>{{end}}
    </div>
  </div>
  {{end}}
  <div class="search-bar">
    <form class="search-form" method="GET" action="/" onsubmit="return false">
      {{if .CurrentOrg}}<input type="hidden" name="org" value="{{.CurrentOrg}}"/>{{end}}
      <input class="search-input" type="text" name="q" id="search-q" value="{{.Search}}"
             placeholder="Поиск по имени, должности, телефону, email..."
             oninput="liveSearch(this.value)" autocomplete="off"/>
      {{if .Search}}<a class="clear-link" href="#" onclick="liveSearch('');document.getElementById('search-q').value='';return false;">&#10005; Сбросить</a>{{end}}
    </form>
  </div>
</div>

<div class="content">
{{range .Departments}}
<div class="dept-card">
  <div class="dept-root-header">
    <span class="dept-root-name">{{.Name}}</span>
  </div>
  {{if .Contacts}}
  <div class="table-wrap">
  <table>
    <thead><tr>
      <th>Каб.</th><th>Должность</th><th>ФИО</th>
      <th>Гор. тел.</th><th>Моб. тел.</th><th>Вн. тел.</th><th>Email</th>
    </tr></thead>
    <tbody>
    {{range .Contacts}}<tr data-s="{{.Room}} {{.Position}} {{.FullName}} {{.PhoneCity}} {{.PhoneMobile}} {{.PhoneInternal}} {{.Email}}">
      <td class="col-room">{{.Room}}</td>
      <td class="col-pos">{{.Position}}</td>
      <td class="col-name">{{.FullName}}</td>
      <td class="col-phone">{{.PhoneCity}}</td>
      <td class="col-phone">{{.PhoneMobile}}</td>
      <td class="col-phone">{{.PhoneInternal}}</td>
      <td class="col-email">{{.Email}}</td>
    </tr>{{end}}
    </tbody>
  </table>
  </div>
  {{end}}
  {{range .Children}}
  <div class="subdept-wrap">
    <div class="subdept-header">
      <span class="subdept-arrow">&#x2514;</span>
      <span class="subdept-name">{{.Name}}</span>
    </div>
    {{if .Contacts}}
    <div class="table-wrap">
    <table>
      <thead><tr>
        <th>Каб.</th><th>Должность</th><th>ФИО</th>
        <th>Гор. тел.</th><th>Моб. тел.</th><th>Вн. тел.</th><th>Email</th>
      </tr></thead>
      <tbody>
      {{range .Contacts}}<tr data-s="{{.Room}} {{.Position}} {{.FullName}} {{.PhoneCity}} {{.PhoneMobile}} {{.PhoneInternal}} {{.Email}}">
        <td class="col-room">{{.Room}}</td>
        <td class="col-pos">{{.Position}}</td>
        <td class="col-name">{{.FullName}}</td>
        <td class="col-phone">{{.PhoneCity}}</td>
        <td class="col-phone">{{.PhoneMobile}}</td>
        <td class="col-phone">{{.PhoneInternal}}</td>
        <td class="col-email">{{.Email}}</td>
      </tr>{{end}}
      </tbody>
    </table>
    </div>
    {{else}}<div class="empty">Нет записей</div>{{end}}
  </div>
  {{end}}
  {{if and (not .Contacts) (not .Children)}}<div class="empty">Нет записей</div>{{end}}
</div>
{{else}}
<div class="no-data">Справочник пуст</div>
{{end}}
<div id="js-no-result" class="no-data" style="display:none">Ничего не найдено</div>
</div>
<script>
(function(){
  function liveSearch(q) {
    q = (q || '').toLowerCase().trim();
    var cards = document.querySelectorAll('.dept-card');
    if (cards.length === 0) return;
    document.querySelectorAll('tr[data-s]').forEach(function(tr){
      tr.hidden = q !== '' && !tr.dataset.s.toLowerCase().includes(q);
    });
    document.querySelectorAll('.subdept-wrap').forEach(function(wrap){
      if (q === '') { wrap.hidden = false; return; }
      wrap.hidden = wrap.querySelectorAll('tr[data-s]:not([hidden])').length === 0;
    });
    var any = false;
    cards.forEach(function(card){
      if (q === '') { card.hidden = false; any = true; return; }
      var vis = card.querySelectorAll('tr[data-s]:not([hidden])').length > 0;
      card.hidden = !vis;
      if (vis) any = true;
    });
    var noRes = document.getElementById('js-no-result');
    if (noRes) noRes.style.display = (!any && q !== '') ? '' : 'none';
  }
  window.liveSearch = liveSearch;
  var initQ = (new URLSearchParams(window.location.search).get('q') || '').trim();
  if (initQ) liveSearch(initQ);
})();
</script>
{{if .Copyright}}<footer class="footer">{{.Copyright}}</footer>{{end}}
{{if .Orgs}}
<script>
(function(){
  const ORGS = {{.OrgsJSON}};
  const input = document.getElementById('org-search');
  const dropdown = document.getElementById('org-dropdown');
  if (!input) return;
  const params = new URLSearchParams(window.location.search);
  const curOrg = parseInt(params.get('org') || '0');
  const curQ   = params.get('q') || '';

  function nav(orgId) {
    const p = new URLSearchParams();
    if (orgId > 0) p.set('org', orgId);
    if (curQ) p.set('q', curQ);
    const qs = p.toString();
    window.location.href = qs ? '/?' + qs : '/';
  }

  function render(list) {
    dropdown.innerHTML = '';
    const all = document.createElement('div');
    all.className = 'org-all-item';
    all.textContent = '— Все организации —';
    all.onclick = () => nav(0);
    dropdown.appendChild(all);
    list.forEach(function(org) {
      const item = document.createElement('div');
      item.className = 'org-item' + (org.id === curOrg ? ' selected' : '');
      item.appendChild(document.createTextNode(org.name));
      if (org.is_default) {
        const b = document.createElement('span');
        b.className = 'org-item-badge';
        b.textContent = 'по умолч.';
        item.appendChild(b);
      }
      item.onclick = () => nav(org.id);
      dropdown.appendChild(item);
    });
    dropdown.style.display = 'block';
  }

  function filtered() {
    const q = input.value.toLowerCase();
    return q ? ORGS.filter(o => o.name.toLowerCase().includes(q)) : ORGS;
  }

  input.removeAttribute('readonly');
  input.addEventListener('focus', function() { render(ORGS); });
  input.addEventListener('input', function() { render(filtered()); });
  input.addEventListener('click',  function() { render(filtered()); });
  document.addEventListener('click', function(e) {
    if (!e.target.closest('.org-selector')) dropdown.style.display = 'none';
  });
})();
</script>
{{end}}
</body>
</html>`

// --- Admin page ---

const adminPage = `<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8"/>
<meta name="viewport" content="width=device-width,initial-scale=1"/>
<title>Администрирование — Телефонный справочник</title>
<style>` + sharedCSS + `</style>
</head>
<body>
<div class="header">
  <div class="header-top">
    <img class="header-logo" src="/logo" onerror="this.style.display='none'" alt=""/>
    <div class="header-title">
      <h1>{{if .SiteTitle}}{{.SiteTitle}}{{else}}Телефонный справочник{{end}} &nbsp;<span class="admin-badge">ADMIN</span></h1>
      {{range .HeaderLines}}<p>{{.}}</p>{{end}}
    </div>
    <div class="header-actions">
      <button class="btn" onclick="openDeptAdd()">+ Подразделение</button>
      <button class="btn-outline" onclick="document.getElementById('dlg-orgs').showModal()">Организации</button>
      <a href="/"><button class="btn-outline" type="button">Просмотр</button></a>
      <form method="POST" action="/logout" style="display:inline">
        <button class="btn-ghost" type="submit">Выйти</button>
      </form>
    </div>
  </div>
  {{if .Orgs}}
  <div class="org-filter">
    <div class="org-selector">
      <input id="org-search-admin" class="org-input" type="text"
             value="{{.OrgName}}"
             placeholder="Фильтр по организации..."
             autocomplete="off" readonly/>
      <div id="org-dropdown-admin" class="org-dropdown" style="display:none"></div>
      {{if .CurrentOrg}}<a class="clear-link" href="/admin?{{if .Search}}q={{.Search}}{{end}}">&#10005; Все орг.</a>{{end}}
    </div>
  </div>
  {{end}}
  <div class="search-bar">
    <form class="search-form" method="GET" action="/admin" onsubmit="return false">
      {{if .CurrentOrg}}<input type="hidden" name="org" value="{{.CurrentOrg}}"/>{{end}}
      <input class="search-input" type="text" name="q" id="search-q" value="{{.Search}}"
             placeholder="Поиск по имени, должности, телефону, email..."
             oninput="liveSearch(this.value)" autocomplete="off"/>
      {{if .Search}}<a class="clear-link" href="#" onclick="liveSearch('');document.getElementById('search-q').value='';return false;">&#10005; Сбросить</a>{{end}}
    </form>
  </div>
</div>

<div class="content">
{{range .Departments}}
<div class="dept-card">
  <div class="dept-root-header">
    <span class="dept-root-name">{{.Name}}</span>
    <div class="dept-root-actions">
      <button class="btn-sm"
        data-id="{{.ID}}" data-name="{{.Name}}" data-sort="{{.SortOrder}}" data-parent="0" data-org="{{.OrganizationID}}"
        onclick="openDeptEdit(this)">&#9998; Изменить</button>
      <button class="btn-sm"
        data-dept="{{.ID}}"
        onclick="openContactAdd(this)">+ Сотрудник</button>
      <button class="btn-sm"
        data-parent="{{.ID}}"
        onclick="openSubDeptAdd(this)">+ Подраздел</button>
      <form method="POST" action="/admin/dept/delete" style="display:inline"
        onsubmit="return confirm('Удалить «{{.Name}}» со всеми подразделами и сотрудниками?')">
        <input type="hidden" name="id" value="{{.ID}}"/>
        <button class="btn-danger-sm" type="submit">&#10005;</button>
      </form>
    </div>
  </div>

  {{if .Contacts}}
  <div class="table-wrap">
  <table>
    <thead><tr>
      <th>Каб.</th><th>Должность</th><th>ФИО</th>
      <th>Гор. тел.</th><th>Моб. тел.</th><th>Вн. тел.</th><th>Email</th><th></th>
    </tr></thead>
    <tbody>
    {{range .Contacts}}<tr data-s="{{.Room}} {{.Position}} {{.FullName}} {{.PhoneCity}} {{.PhoneMobile}} {{.PhoneInternal}} {{.Email}}">
      <td class="col-room">{{.Room}}</td>
      <td class="col-pos">{{.Position}}</td>
      <td class="col-name">{{.FullName}}</td>
      <td class="col-phone">{{.PhoneCity}}</td>
      <td class="col-phone">{{.PhoneMobile}}</td>
      <td class="col-phone">{{.PhoneInternal}}</td>
      <td class="col-email">{{.Email}}</td>
      <td class="col-act"><div class="act-row">
        <button class="btn-sm"
          data-id="{{.ID}}" data-dept="{{.DepartmentID}}"
          data-room="{{.Room}}" data-position="{{.Position}}" data-fullname="{{.FullName}}"
          data-city="{{.PhoneCity}}" data-mobile="{{.PhoneMobile}}"
          data-internal="{{.PhoneInternal}}" data-email="{{.Email}}"
          onclick="openContactEdit(this)">&#9998;</button>
        <form method="POST" action="/admin/contact/delete" onsubmit="return confirm('Удалить запись?')">
          <input type="hidden" name="id" value="{{.ID}}"/>
          <button class="btn-danger-sm" type="submit">&#10005;</button>
        </form>
      </div></td>
    </tr>{{end}}
    </tbody>
  </table>
  </div>
  {{end}}

  {{range .Children}}
  <div class="subdept-wrap">
    <div class="subdept-header">
      <span class="subdept-arrow">&#x2514;</span>
      <span class="subdept-name">{{.Name}}</span>
      <div class="subdept-actions">
        <button class="btn-sm"
          data-id="{{.ID}}" data-name="{{.Name}}" data-sort="{{.SortOrder}}" data-parent="{{.ParentID}}" data-org="{{.OrganizationID}}"
          onclick="openDeptEdit(this)">&#9998; Изменить</button>
        <button class="btn-sm"
          data-dept="{{.ID}}"
          onclick="openContactAdd(this)">+ Сотрудник</button>
        <form method="POST" action="/admin/dept/delete" style="display:inline"
          onsubmit="return confirm('Удалить «{{.Name}}» со всеми сотрудниками?')">
          <input type="hidden" name="id" value="{{.ID}}"/>
          <button class="btn-danger-sm" type="submit">&#10005;</button>
        </form>
      </div>
    </div>
    {{if .Contacts}}
    <div class="table-wrap">
    <table>
      <thead><tr>
        <th>Каб.</th><th>Должность</th><th>ФИО</th>
        <th>Гор. тел.</th><th>Моб. тел.</th><th>Вн. тел.</th><th>Email</th><th></th>
      </tr></thead>
      <tbody>
      {{range .Contacts}}<tr data-s="{{.Room}} {{.Position}} {{.FullName}} {{.PhoneCity}} {{.PhoneMobile}} {{.PhoneInternal}} {{.Email}}">
        <td class="col-room">{{.Room}}</td>
        <td class="col-pos">{{.Position}}</td>
        <td class="col-name">{{.FullName}}</td>
        <td class="col-phone">{{.PhoneCity}}</td>
        <td class="col-phone">{{.PhoneMobile}}</td>
        <td class="col-phone">{{.PhoneInternal}}</td>
        <td class="col-email">{{.Email}}</td>
        <td class="col-act"><div class="act-row">
          <button class="btn-sm"
            data-id="{{.ID}}" data-dept="{{.DepartmentID}}"
            data-room="{{.Room}}" data-position="{{.Position}}" data-fullname="{{.FullName}}"
            data-city="{{.PhoneCity}}" data-mobile="{{.PhoneMobile}}"
            data-internal="{{.PhoneInternal}}" data-email="{{.Email}}"
            onclick="openContactEdit(this)">&#9998;</button>
          <form method="POST" action="/admin/contact/delete" onsubmit="return confirm('Удалить запись?')">
            <input type="hidden" name="id" value="{{.ID}}"/>
            <button class="btn-danger-sm" type="submit">&#10005;</button>
          </form>
        </div></td>
      </tr>{{end}}
      </tbody>
    </table>
    </div>
    {{else}}<div class="empty">Нет записей</div>{{end}}
  </div>
  {{end}}

  {{if and (not .Contacts) (not .Children)}}<div class="empty">Нет записей</div>{{end}}
</div>
{{else}}
<div class="no-data">Нет подразделений. <button class="btn" onclick="openDeptAdd()">Добавить первое</button></div>
{{end}}
<div id="js-no-result" class="no-data" style="display:none">Ничего не найдено</div>
</div>

<!-- Добавить подразделение -->
<dialog id="dlg-dept-add">
  <div class="dlg-head"><h3>Добавить подразделение</h3></div>
  <form method="POST" action="/admin/dept/add">
    <input type="hidden" name="parent_id" id="da-parent" value="0"/>
    <div class="dlg-body">
      <label>Название
        <input type="text" name="name" id="da-name" required autofocus placeholder="например: АППАРАТ ДУМЫ"/>
      </label>
      <div class="row2">
        <label>Родительское подразделение
          <select name="_parent_display" id="da-parent-sel" onchange="document.getElementById('da-parent').value=this.value">
            <option value="0">— корневое —</option>
            {{range .RootDepts}}<option value="{{.ID}}">{{.Name}}</option>{{end}}
          </select>
        </label>
        <label>Порядок сортировки
          <input type="number" name="sort_order" value="0"/>
        </label>
      </div>
      {{if .Orgs}}
      <label>Организация
        <select name="organization_id" id="da-org">
          <option value="0">— не указана —</option>
          {{range .Orgs}}<option value="{{.ID}}">{{.Name}}</option>{{end}}
        </select>
      </label>
      {{end}}
    </div>
    <div class="dlg-foot">
      <button class="btn-ghost" type="button" onclick="this.closest('dialog').close()">Отмена</button>
      <button class="btn" type="submit">Сохранить</button>
    </div>
  </form>
</dialog>

<!-- Изменить подразделение -->
<dialog id="dlg-dept-edit">
  <div class="dlg-head"><h3>Изменить подразделение</h3></div>
  <form method="POST" action="/admin/dept/edit">
    <input type="hidden" name="id" id="ed-id"/>
    <input type="hidden" name="parent_id" id="ed-parent"/>
    <div class="dlg-body">
      <label>Название
        <input type="text" name="name" id="ed-name" required/>
      </label>
      <div class="row2">
        <label>Родительское подразделение
          <select id="ed-parent-sel" onchange="document.getElementById('ed-parent').value=this.value">
            <option value="0">— корневое —</option>
            {{range .RootDepts}}<option value="{{.ID}}">{{.Name}}</option>{{end}}
          </select>
        </label>
        <label>Порядок сортировки
          <input type="number" name="sort_order" id="ed-sort"/>
        </label>
      </div>
      {{if .Orgs}}
      <label>Организация
        <select name="organization_id" id="ed-org">
          <option value="0">— не указана —</option>
          {{range .Orgs}}<option value="{{.ID}}">{{.Name}}</option>{{end}}
        </select>
      </label>
      {{end}}
    </div>
    <div class="dlg-foot">
      <button class="btn-ghost" type="button" onclick="this.closest('dialog').close()">Отмена</button>
      <button class="btn" type="submit">Сохранить</button>
    </div>
  </form>
</dialog>

<!-- Добавить сотрудника -->
<dialog id="dlg-contact-add">
  <div class="dlg-head"><h3>Добавить сотрудника</h3></div>
  <form method="POST" action="/admin/contact/add">
    <input type="hidden" name="dept_id" id="ac-dept"/>
    <div class="dlg-body">
      <div class="row2">
        <label>Организация
          <select id="ac-org-filter" onchange="acFilterDepts(this.value)">
            <option value="0">— Все организации —</option>
          </select>
        </label>
        <label>Подразделение <span style="color:#e03131">*</span>
          <select id="ac-dept-sel" required onchange="document.getElementById('ac-dept').value=this.value">
            <option value="">— выберите —</option>
          </select>
        </label>
      </div>
      <label>Должность
        <input type="text" name="position" placeholder="Начальник отдела..."/>
      </label>
      <label>ФИО
        <input type="text" name="full_name" placeholder="Иванов Иван Иванович"/>
      </label>
      <div class="row3">
        <label>Каб. <input type="text" name="room"/></label>
        <label>Вн. тел. <input type="text" name="phone_internal"/></label>
        <label>Гор. тел. <input type="text" name="phone_city"/></label>
      </div>
      <div class="row2">
        <label>Моб. тел. <input type="text" name="phone_mobile"/></label>
        <label>Email <input type="text" name="email"/></label>
      </div>
    </div>
    <div class="dlg-foot">
      <button class="btn-ghost" type="button" onclick="this.closest('dialog').close()">Отмена</button>
      <button class="btn" type="submit">Сохранить</button>
    </div>
  </form>
</dialog>

<!-- Изменить сотрудника -->
<dialog id="dlg-contact-edit">
  <div class="dlg-head"><h3>Изменить сотрудника</h3></div>
  <form method="POST" action="/admin/contact/edit">
    <input type="hidden" name="id" id="ec-id"/>
    <input type="hidden" name="dept_id" id="ec-dept"/>
    <div class="dlg-body">
      <label>Должность
        <input type="text" name="position" id="ec-pos"/>
      </label>
      <label>ФИО
        <input type="text" name="full_name" id="ec-name"/>
      </label>
      <div class="row3">
        <label>Каб. <input type="text" name="room" id="ec-room"/></label>
        <label>Вн. тел. <input type="text" name="phone_internal" id="ec-int"/></label>
        <label>Гор. тел. <input type="text" name="phone_city" id="ec-city"/></label>
      </div>
      <div class="row2">
        <label>Моб. тел. <input type="text" name="phone_mobile" id="ec-mob"/></label>
        <label>Email <input type="text" name="email" id="ec-email"/></label>
      </div>
    </div>
    <div class="dlg-foot">
      <button class="btn-ghost" type="button" onclick="this.closest('dialog').close()">Отмена</button>
      <button class="btn" type="submit">Сохранить</button>
    </div>
  </form>
</dialog>

<!-- Управление организациями -->
<dialog id="dlg-orgs">
  <div class="dlg-head"><h3>Организации</h3></div>
  <div class="dlg-body">
    {{if .Orgs}}
    <div class="table-wrap">
    <table>
      <thead><tr><th>Название</th><th style="text-align:center">Умолч.</th><th></th></tr></thead>
      <tbody>
      {{range .Orgs}}<tr>
        <td style="font-size:13px;font-weight:500">{{.Name}}</td>
        <td style="text-align:center;color:#3b5bdb;font-size:14px">{{if .IsDefault}}&#10003;{{end}}</td>
        <td><div class="act-row">
          <form method="POST" action="/admin/org/setdefault" style="display:inline">
            <input type="hidden" name="id" value="{{.ID}}"/>
            <button class="btn-sm" type="submit" title="Сделать организацией по умолчанию">&#9733; Умолч.</button>
          </form>
          <button class="btn-sm" data-id="{{.ID}}" data-name="{{.Name}}" onclick="openOrgEdit(this)">&#9998;</button>
          <form method="POST" action="/admin/org/delete" style="display:inline"
            onsubmit="return confirm('Удалить организацию «{{.Name}}»? Подразделения останутся, но потеряют привязку.')">
            <input type="hidden" name="id" value="{{.ID}}"/>
            <button class="btn-danger-sm" type="submit">&#10005;</button>
          </form>
        </div></td>
      </tr>{{end}}
      </tbody>
    </table>
    </div>
    {{else}}
    <div class="empty">Нет организаций</div>
    {{end}}
    <div style="border-top:1px solid #edf0f9;padding-top:14px;margin-top:14px">
      <form method="POST" action="/admin/org/add">
        <div style="display:flex;gap:8px">
          <input type="text" name="name" placeholder="Название новой организации" required style="flex:1"/>
          <button class="btn" type="submit">Добавить</button>
        </div>
      </form>
    </div>
  </div>
  <div class="dlg-foot">
    <button class="btn-ghost" type="button" onclick="this.closest('dialog').close()">Закрыть</button>
  </div>
</dialog>

<!-- Изменить организацию -->
<dialog id="dlg-org-edit">
  <div class="dlg-head"><h3>Изменить организацию</h3></div>
  <form method="POST" action="/admin/org/edit">
    <input type="hidden" name="id" id="oe-id"/>
    <div class="dlg-body">
      <label>Название
        <input type="text" name="name" id="oe-name" required/>
      </label>
    </div>
    <div class="dlg-foot">
      <button class="btn-ghost" type="button" onclick="this.closest('dialog').close()">Отмена</button>
      <button class="btn" type="submit">Сохранить</button>
    </div>
  </form>
</dialog>

<script>
var DEPT_OPTS = {{.DeptOptsJSON}};
var ADMIN_ORGS = {{.OrgsJSON}};

// Живой поиск
(function(){
  function liveSearch(q) {
    q = (q || '').toLowerCase().trim();
    var cards = document.querySelectorAll('.dept-card');
    if (cards.length === 0) return;
    document.querySelectorAll('tr[data-s]').forEach(function(tr){
      tr.hidden = q !== '' && !tr.dataset.s.toLowerCase().includes(q);
    });
    document.querySelectorAll('.subdept-wrap').forEach(function(wrap){
      if (q === '') { wrap.hidden = false; return; }
      wrap.hidden = wrap.querySelectorAll('tr[data-s]:not([hidden])').length === 0;
    });
    var any = false;
    cards.forEach(function(card){
      if (q === '') { card.hidden = false; any = true; return; }
      var vis = card.querySelectorAll('tr[data-s]:not([hidden])').length > 0;
      card.hidden = !vis;
      if (vis) any = true;
    });
    var noRes = document.getElementById('js-no-result');
    if (noRes) noRes.style.display = (!any && q !== '') ? '' : 'none';
  }
  window.liveSearch = liveSearch;
  var initQ = (new URLSearchParams(window.location.search).get('q') || '').trim();
  if (initQ) liveSearch(initQ);
})();

// Org dropdown для фильтра на admin-странице
(function(){
  var input = document.getElementById('org-search-admin');
  var dropdown = document.getElementById('org-dropdown-admin');
  if (!input || !dropdown) return;
  var params = new URLSearchParams(window.location.search);
  var curOrg = parseInt(params.get('org') || '0');
  var curQ   = params.get('q') || '';

  function nav(orgId) {
    var p = new URLSearchParams();
    if (orgId > 0) p.set('org', orgId);
    if (curQ) p.set('q', curQ);
    var qs = p.toString();
    window.location.href = qs ? '/admin?' + qs : '/admin';
  }

  function render(list) {
    dropdown.innerHTML = '';
    var all = document.createElement('div');
    all.className = 'org-all-item';
    all.textContent = '— Все организации —';
    all.onclick = function(){ nav(0); };
    dropdown.appendChild(all);
    list.forEach(function(org) {
      var item = document.createElement('div');
      item.className = 'org-item' + (org.id === curOrg ? ' selected' : '');
      item.appendChild(document.createTextNode(org.name));
      if (org.is_default) {
        var b = document.createElement('span');
        b.className = 'org-item-badge';
        b.textContent = 'по умолч.';
        item.appendChild(b);
      }
      item.onclick = function(){ nav(org.id); };
      dropdown.appendChild(item);
    });
    dropdown.style.display = 'block';
  }

  function filtered() {
    var q = input.value.toLowerCase();
    return q ? ADMIN_ORGS.filter(function(o){ return o.name.toLowerCase().includes(q); }) : ADMIN_ORGS;
  }

  input.removeAttribute('readonly');
  input.addEventListener('focus', function(){ render(ADMIN_ORGS); });
  input.addEventListener('input', function(){ render(filtered()); });
  input.addEventListener('click',  function(){ render(filtered()); });
  document.addEventListener('click', function(e) {
    if (!e.target.closest('.org-selector')) dropdown.style.display = 'none';
  });
})();

// Заполняем фильтр организаций в диалоге добавления сотрудника
(function() {
  var orgSel = document.getElementById('ac-org-filter');
  if (!orgSel) return;
  var seen = {};
  DEPT_OPTS.forEach(function(d) {
    if (!seen[d.org_id]) {
      seen[d.org_id] = true;
      var opt = document.createElement('option');
      opt.value = d.org_id;
      opt.textContent = d.org_name;
      orgSel.appendChild(opt);
    }
  });
  acFilterDepts(0);
})();

function acFilterDepts(orgId) {
  orgId = parseInt(orgId);
  var sel = document.getElementById('ac-dept-sel');
  sel.innerHTML = '<option value="">— выберите —</option>';
  DEPT_OPTS.forEach(function(d) {
    if (orgId === 0 || d.org_id === orgId) {
      var opt = document.createElement('option');
      opt.value = d.id;
      opt.textContent = d.name;
      sel.appendChild(opt);
    }
  });
  document.getElementById('ac-dept').value = '';
}

function openDeptAdd() {
  var curOrg = {{.CurrentOrg}};
  document.getElementById('da-name').value = '';
  document.getElementById('da-parent').value = '0';
  var parSel = document.getElementById('da-parent-sel');
  if (parSel) parSel.value = '0';
  var orgSel = document.getElementById('da-org');
  if (orgSel && curOrg > 0) orgSel.value = curOrg;
  document.getElementById('dlg-dept-add').showModal();
}
function openDeptEdit(btn) {
  const d = btn.dataset;
  document.getElementById('ed-id').value   = d.id;
  document.getElementById('ed-name').value = d.name;
  document.getElementById('ed-sort').value = d.sort;
  document.getElementById('ed-parent').value = d.parent;
  const sel = document.getElementById('ed-parent-sel');
  for (let opt of sel.options) opt.selected = (opt.value == d.parent);
  const orgSel = document.getElementById('ed-org');
  if (orgSel) for (let opt of orgSel.options) opt.selected = (opt.value == (d.org || '0'));
  document.getElementById('dlg-dept-edit').showModal();
}
function openOrgEdit(btn) {
  document.getElementById('oe-id').value   = btn.dataset.id;
  document.getElementById('oe-name').value = btn.dataset.name;
  document.getElementById('dlg-org-edit').showModal();
}
function openSubDeptAdd(btn) {
  const parentID = btn.dataset.parent;
  const sel = document.getElementById('da-parent-sel');
  for (let opt of sel.options) opt.selected = (opt.value == parentID);
  document.getElementById('da-parent').value = parentID;
  document.getElementById('da-name').value = '';
  document.getElementById('dlg-dept-add').showModal();
}
function openContactAdd(btn) {
  const deptId = parseInt(btn.dataset.dept);
  // Находим организацию этого отдела
  const dept = DEPT_OPTS.find(function(d){ return d.id === deptId; });
  const orgId = dept ? dept.org_id : 0;
  // Устанавливаем фильтр организации и список отделов
  const orgSel = document.getElementById('ac-org-filter');
  if (orgSel) { orgSel.value = orgId; acFilterDepts(orgId); }
  // Выбираем нужный отдел
  document.getElementById('ac-dept').value = deptId;
  const deptSel = document.getElementById('ac-dept-sel');
  if (deptSel) deptSel.value = deptId;
  document.getElementById('dlg-contact-add').showModal();
}
function openContactEdit(btn) {
  const d = btn.dataset;
  document.getElementById('ec-id').value    = d.id;
  document.getElementById('ec-dept').value  = d.dept;
  document.getElementById('ec-room').value  = d.room;
  document.getElementById('ec-pos').value   = d.position;
  document.getElementById('ec-name').value  = d.fullname;
  document.getElementById('ec-city').value  = d.city;
  document.getElementById('ec-mob').value   = d.mobile;
  document.getElementById('ec-int').value   = d.internal;
  document.getElementById('ec-email').value = d.email;
  document.getElementById('dlg-contact-edit').showModal();
}
</script>
{{if .Copyright}}<footer class="footer">{{.Copyright}}</footer>{{end}}
</body>
</html>`
