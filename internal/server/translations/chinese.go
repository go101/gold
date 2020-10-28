package translations

import (
	"fmt"
	"time"

	"go101.org/gold/code"
)

type Chinese struct{}

func (*Chinese) Name() string { return "简体中文" }

func (*Chinese) LangTag() string { return "zh-CN" }

///////////////////////////////////////////////////////////////////
// common
///////////////////////////////////////////////////////////////////

func (*Chinese) Text_Space() string { return "" }

func (*Chinese) Text_Comma() string { return "，" }

func (*Chinese) Text_Colon(tailSpace bool) string { return "：" }

func (*Chinese) Text_Period(paragraphEnd bool) string { return "。" }

func (*Chinese) Text_EnclosedInOarentheses(text string) string {
	return "（" + text + "）"
}

func (*Chinese) Text_PreferredFontList() string {
	return `"Courier New", Courier, monospace, "Microsoft YaHei", "宋体"`
}

func (*Chinese) Text_BlankList() string {
	return "（无）"
}

///////////////////////////////////////////////////////////////////
// server
///////////////////////////////////////////////////////////////////

func (*Chinese) Text_Server_Started() string {
	return "服务已启动："
}

///////////////////////////////////////////////////////////////////
// analyzing
///////////////////////////////////////////////////////////////////

func (*Chinese) Text_Analyzing() string { return "分析中......" }

func (*Chinese) Text_AnalyzingRefresh(currentPageURL string) string {
	return fmt.Sprintf(`正在分析，稍等片刻......（<a href="%s">刷新</a>）`, currentPageURL)
}

func (*Chinese) Text_Analyzing_Start() string {
	return "开始分析......"
}

func (*Chinese) Text_Analyzing_PreparationDone(d time.Duration) string {
	return fmt.Sprintf("准备完毕：%s", d)
}

func (*Chinese) Text_Analyzing_NFilesParsed(numFiles int, d time.Duration) string {
	return fmt.Sprintf("解析了%d个源文件：%s", numFiles, d)
}

func (*Chinese) Text_Analyzing_ParsePackagesDone(numFiles int, d time.Duration) string {
	return fmt.Sprintf("全部%d个源文件解析完毕：%s", numFiles, d)
}

func (*Chinese) Text_Analyzing_CollectPackages(numPkgs int, d time.Duration) string {
	return fmt.Sprintf("搜集了%d个代码包：%s", numPkgs, d)
}

func (*Chinese) Text_Analyzing_SortPackagesByDependencies(d time.Duration) string {
	return fmt.Sprintf("按依赖关系对代码包进行排序：%s", d)
}

func (*Chinese) Text_Analyzing_CollectDeclarations(d time.Duration) string {
	return fmt.Sprintf("搜集各种声明：%s", d)
}

func (*Chinese) Text_Analyzing_CollectRuntimeFunctionPositions(d time.Duration) string {
	return fmt.Sprintf("搜集一些runtime包中的函数的代码位置：%s", d)
}

func (*Chinese) Text_Analyzing_FindTypeSources(d time.Duration) string {
	return fmt.Sprintf("寻找类型声明的源类型：%s", d)
}

func (*Chinese) Text_Analyzing_CollectSelectors(d time.Duration) string {
	return fmt.Sprintf("搜集选择器：%s", d)
}

func (*Chinese) Text_Analyzing_FindImplementations(d time.Duration) string {
	return fmt.Sprintf("寻找类型实现关系：%s", d)
}

func (*Chinese) Text_Analyzing_RegisterInterfaceMethodsForTypes(d time.Duration) string {
	return fmt.Sprintf("注册接口方法：%s", d)
}

func (*Chinese) Text_Analyzing_MakeStatistics(d time.Duration) string {
	return fmt.Sprintf("整理统计：%s", d)
}

func (*Chinese) Text_Analyzing_CollectSourceFiles(d time.Duration) string {
	return fmt.Sprintf("搜集源文件：%s", d)
}

func (*Chinese) Text_Analyzing_CollectObjectReferences(d time.Duration) string {
	return fmt.Sprintf("搜集代码元素对象引用：%s", d)
}

func (*Chinese) Text_Analyzing_CacheSourceFiles(d time.Duration) string {
	return fmt.Sprintf("缓存源文件：%s", d)
}

func (*Chinese) Text_Analyzing_Done(d time.Duration, memoryUse string) string {
	return fmt.Sprintf("分析完毕（共用时%s，最终消耗内存%s）", d, memoryUse)
}

///////////////////////////////////////////////////////////////////
// overview page
///////////////////////////////////////////////////////////////////

func (*Chinese) Text_Overview() string { return "概览" }

func (*Chinese) Text_PackageList() string {
	return "代码包列表"
}

func (*Chinese) Text_StatisticsWithMoreLink(detailedStatsLink string) string {
	return fmt.Sprintf(`统计信息（<a href="%s">更多详细信息</a>）`, detailedStatsLink)
}

func (*Chinese) Text_SimpleStats(stats *code.Stats) string {
	return fmt.Sprintf(`分析了%d个代码包，解析了%d个Go源文件。
平均说来：
* 每个Go源文件引入了个%.2f代码包；
* 每个代码包依赖于%.2f个其它代码包，
  含有%.2f个源文件，并且导出了
  - %.2f个类型名；
  - %.2f个变量；
  - %.2f个常量；
  - %.2f个函数。`,
		stats.Packages, stats.AstFiles,
		float64(stats.Imports)/float64(stats.AstFiles),
		float64(stats.AllPackageDeps)/float64(stats.Packages),
		float64(stats.FilesWithoutGenerateds)/float64(stats.Packages),
		float64(stats.ExportedTypeNames)/float64(stats.Packages),
		float64(stats.ExportedVariables)/float64(stats.Packages),
		float64(stats.ExportedConstants)/float64(stats.Packages),
		float64(stats.ExportedFunctions)/float64(stats.Packages),
	)

}

func (*Chinese) Text_Modules() string { return "模块列表" }

func (*Chinese) Text_BelongingModule() string { return "所属模块" }

func (*Chinese) Text_RequireStat(numRequires, numRequiredBys int) string {
	return fmt.Sprintf("需要%d模块，并且被%d个模块所需要。", numRequires, numRequiredBys)
}

func (*Chinese) Text_UpdateTip(tipName string) string {
	switch tipName {
	case "ToUpdate":
		return `<b>Golds</b>已经有一个多月没有更新了，运行<b>go get -u go101.org/golds</b>或者<b><a href="/update">点击</a></b>来更新它。`
	case "Updating":
		return `<b>Golds</b>正在被更新中.....`
	case "Updated":
		return `<b>Golds</b>已经被更新了，重启此golds服务进程可以看到最新的效果。`
	}
	return ""
}

func (*Chinese) Text_SortBy() string {
	return "排序依据："
}

func (*Chinese) Text_Filter() string {
	return "列出："
}

func (*Chinese) Text_SortByItem(by string) string {
	switch by {
	case "alphabet":
		return "按字母排序"
	case "popularity":
		return "按流行度排序"
	case "importedbys":
		return "按被引入量排序"
	default:
		panic("unknown sort-by: " + by)
	}
}

func (*Chinese) Text_FilterItem(fltr string) string {
	switch fltr {
	case "all":
		return "所有"
	case "mainpackages":
		return "main包"
	case "testingpackages":
		return "测试包"
	default:
		panic("unknown filter item: " + fltr)
	}
}

///////////////////////////////////////////////////////////////////
// package details page: type details
///////////////////////////////////////////////////////////////////

func (*Chinese) Text_Package(pkgPath string) string {
	return fmt.Sprintf("代码包：%s", pkgPath)
}

func (*Chinese) Text_BelongingPackage() string { return "所属代码包" }

func (*Chinese) Text_PackageDocsLinksOnOtherWebsites(pkgPath string, isStdPkg bool) string {
	if isStdPkg {
		return fmt.Sprintf(`<i> （在 <a href="https://golang.google.cn/pkg/%[1]s/" target="_blank">golang.google.cn</a> 和 <a href="https://pkg.go.dev/%[1]s" target="_blank">go.dev</a> 上）</i>`, pkgPath)
	} else {
		return fmt.Sprintf(`<i> （在 <a href="https://pkg.go.dev/%s" target="_blank">go.dev</a> 上）</i>`, pkgPath)
	}
}

func (*Chinese) Text_ImportPath() string { return "引入路径" }

func (*Chinese) Text_ImportStat(numImports, numImportedBys int, depPageURL string) string {
	importsStr := fmt.Sprintf("%d个代码包", numImports)
	if numImports > 0 {
		importsStr = fmt.Sprintf(`<a href="%s">%s</a>`, depPageURL, importsStr)
	}

	importedBysStr := fmt.Sprintf("%d个代码包", numImportedBys)
	if numImportedBys > 0 {
		importedBysStr = fmt.Sprintf(`<a href="%s#imported-by">%s</a>`, depPageURL, importedBysStr)
	}

	return fmt.Sprintf(`引入了%s，并被%s引入。`, importsStr, importedBysStr)
}

func (*Chinese) Text_InvolvedFiles(num int) string { return "相关源文件" }

func (*Chinese) Text_ExportedValues(num int) string {
	return "导出值"
}

func (*Chinese) Text_ExportedTypeNames(num int) string {
	return "导出类型名"
}

func (*Chinese) Text_AllPackageLevelTypeNames(num int) string {
	return "包级类型名"
}

func (*Chinese) Text_TypeNameListShowOption(exportedsOnly bool) string {
	if exportedsOnly {
		return "只展示导出类型"
	} else {
		return "亦展示非导出类型"
	}
}

///////////////////////////////////////////////////////////////////
// package details page: type details
///////////////////////////////////////////////////////////////////

func (*Chinese) Text_Fields(num int, exportedsOnly bool) string {
	if exportedsOnly {
		return fmt.Sprintf("%d个导出字段", num)
	} else {
		return fmt.Sprintf("%d个字段", num)
	}
}

func (*Chinese) Text_Methods(num int, exportedsOnly bool) string {
	if exportedsOnly {
		return fmt.Sprintf("%d个导出方法", num)
	} else {
		return fmt.Sprintf("%d个方法", num)
	}
}

func (*Chinese) Text_ImplementedBy(num int) string {
	return fmt.Sprintf("被%d+类型实现", num)
}

func (*Chinese) Text_Implements(num int) string {
	return fmt.Sprintf("实现了%d+接口", num)
}

func (*Chinese) Text_AsOutputsOf(num int) string {
	return fmt.Sprintf("使用此类型做为结果的函数（%d+）", num)
}

func (*Chinese) Text_AsInputsOf(num int) string {
	return fmt.Sprintf("使用此类型做为参数的函数（%d+）", num)
}

func (*Chinese) Text_AsTypesOf(num int) string {
	return fmt.Sprintf("和此类型相关的值（%d+）", num)
}

func (*Chinese) Text_References(num int) string {
	return fmt.Sprintf("引用（%d+）", num)
}

///////////////////////////////////////////////////////////////////
// package dependencies page
///////////////////////////////////////////////////////////////////

func (*Chinese) Text_DependencyRelations(pkgPath string) string {
	if pkgPath == "" {
		return "依赖关系"
	} else {
		return fmt.Sprintf("依赖关系：%s", pkgPath)
	}
}

func (*Chinese) Text_Imports() string { return "引入了这些代码包" }

func (*Chinese) Text_ImportedBy() string { return "被这些代码包引入" }

///////////////////////////////////////////////////////////////////
// method implementation page
///////////////////////////////////////////////////////////////////

func (*Chinese) Text_MethodImplementations() string {
	return "方法实现列表"
}

func (*Chinese) Text_NumMethodsImplementingNothing(count int) string {
	if count == 0 {
		return ""
	}
	return fmt.Sprintf("（%d个其它方法什么也没实现）", count)
}

///////////////////////////////////////////////////////////////////
// object references(uses) page
///////////////////////////////////////////////////////////////////

func (*Chinese) Text_ReferenceList() string {
	return "引用列表"
}

func (*Chinese) Text_ObjectKind(kind string) string {
	switch kind {
	case "field":
		return "字段"
	case "method":
		return "方法"
	default:
		panic("unknown object kind name: " + kind)
	}
}

func (*Chinese) Text_ObjectUses(num int) string {
	return fmt.Sprintf("%d处使用", num)
}

///////////////////////////////////////////////////////////////////
// source code page
///////////////////////////////////////////////////////////////////

func (*Chinese) Text_SourceCode(pkgPath, bareFilename string) string {
	return fmt.Sprintf("源文件：%s（%s代码包中）", bareFilename, pkgPath)
}

func (*Chinese) Text_SourceFilePath() string { return "源文件：" }

func (*Chinese) Text_GeneratedFrom() string { return "从此文件生成" }

///////////////////////////////////////////////////////////////////
// statistics
///////////////////////////////////////////////////////////////////

func (*Chinese) Text_Statistics() string {
	return "统计信息"
}

func (*Chinese) Text_ChartTitle(chartName string) string {
	switch chartName {
	case "gosourcefiles-by-imports":
		return "Go源文件数量按照引入数量的分布"
	case "packages-by-dependencies":
		return "库包数量按照依赖数量的分布"
	case "exportedtypenames-by-kinds":
		return "导出的类型名数量按照类型种类的分布"
	case "exportedstructtypes-by-embeddingfields":
		return "导出的结构体类型数量按照内嵌字段数量的分布"
	//case "exportedstructtypes-by-allfields":
	//	return "导出的结构体类型数量按照字段数量的分布"
	case "exportedstructtypes-by-explicitfields":
		return "导出的结构体类型数量按照显式字段数量的分布"
	//case "exportedstructtypes-by-exportedfields":
	//	return "导出的结构体类型数量按照导出字段数量的分布"
	case "exportedstructtypes-by-exportedexplicitfields":
		return "导出的结构体类型数量按照导出显式字段数量的分布"
	case "exportedstructtypes-by-exportedpromotedfields":
		return "导出的结构体类型数量按照导出提升字段数量的分布"
	case "exportedfunctions-by-parameters":
		return "导出的函数（包括方法）数量按照参数个数的分布"
	case "exportedfunctions-by-results":
		return "导出的函数（包括方法）数量按照返回结果个数的分布"
	case "exportedidentifiers-by-lengths":
		return "导出的标识符数量按照标识符长度的分布"
	case "exportedvariables-by-typekinds":
		return "导出的变量数量按照变量类型种类的分布"
	case "exportedconstants-by-typekinds":
		return "导出的常量数量按照常量类型（或者默认类型）种类的分布"
	case "exportednoninterfacetypes-by-exportedmethods":
		return "导出的非接口类型名数量按照导出方法数的分布"
	case "exportedinterfacetypes-by-exportedmethods":
		return "导出的接口类型名数量按照导出方法数的分布"
	default:
		panic("unknown char name: " + chartName)
	}
}

func (*Chinese) Text_StatisticsTitle(titleName string) string {
	switch titleName {
	case "packages":
		return "库包"
	case "types":
		return "类型"
	case "values":
		return "值（变量/常量/函数）"
	case "others":
		return "其它"
	default:
		panic("unknown statistics tile: " + titleName)
	}
}

func (*Chinese) Text_PackageStatistics(values map[string]interface{}) string {
	return fmt.Sprintf(`
	共<a href="%s">%d个库包</a>，其中%d个是标准库包。
	共%d个源文件，其中%d个为Go源文件。
	平均说来：
	- 每个库包含有%.2f个源文件；
	- 每个Go源文件引入了%.2f个库包；
	- 每个库包依赖于%.2f个其它库包。

	<img src="%s"></image>
	<img src="%s"></image>
`,
		values["overviewPageURL"],
		values["packageCount"],
		values["standardPackageCount"],
		values["sourceFileCount"],
		values["goSourceFileCount"],
		values["averageSourceFileCountPerPackage"],
		values["averageImportCountPerFile"],
		values["averageDependencyCountPerPackage"],

		values["gosourcefilesByImportsChartURL"],
		values["packagesByDependenciesChartURL"],
	)
}

func (*Chinese) Text_TypeStatistics(values map[string]interface{}) string {
	return fmt.Sprintf(`
	共%d个导出类型名，其中%d个为类型别名。
	它们中有%d个为组合类型、%d个为基本类型。
	在基本类型中，%d个为整数型（其中%d个为无符号类型）。

	<img src="%s"></image>

	在%d个导出结构体类型中，%d个含有内嵌字段，%d个拥有提升字段。

	<img src="%s"></image>

	平均说来，每个导出结构体类型拥有
	* %.2f个字段（包括提升字段和非导出字段）；
	* %.2f个显式字段（包括非导出字段）；
	* %.2f个导出字段（包括提升字段）；
	* %.2f个导出显式字段。

	<img src="%s"></image>
	<img src="%s"></image>
	<img src="%s"></image>

	平均说来，
	- 对于拥有至少一个导出方法的导出非接口类型，每个拥有%.2f个导出方法。
	- 每个导出接口类型指定了%.2f个导出方法。

	<img src="%s"></image>
	<img src="%s"></image>
`,
		values["exportedTypeNameCount"],
		values["exportedTypeAliases"],
		values["exportedCompositeTypeNames"],
		values["exportedBasicTypeNames"],
		values["exportedIntergerTypeNames"],
		values["exportedUnsignedTypeNames"],

		values["exportedtypenamesByKindsChartURL"],

		values["exportedStructTypeNames"],
		values["exportedNamedStructTypesWithEmbeddingFields"],
		values["exportedNamedStructTypesWithPromotedFields"],

		values["exportedstructtypesByEmbeddingfieldsChartURL"],

		values["exportedNamedStructTypeFieldsPerExportedStruct"],
		values["exportedNamedStructTypeExplicitFieldsPerExportedStruct"],
		values["exportedNamedStructTypeExportedFieldsPerExportedStruct"],
		values["exportedNamedStructTypeExportedExplicitFieldsPerExportedStruct"],

		values["exportedstructtypesByExplicitfieldsChartURL"],
		values["exportedstructtypesByExportedexplicitfieldsChartURL"],
		//values["exportedstructtypesByExportedfieldsChartURL"],
		values["exportedstructtypesByExportedpromotedfieldsChartURL"],

		values["exportedNamedNonInterfacesExportedMethodsPerExportedNonInterfaceType"],
		values["exportedNamedInterfacesExportedMethodsPerExportedInterfaceType"],

		values["exportednoninterfacetypesByExportedmethodsChartURL"],
		values["exportedinterfacetypesByExportedmethodsChartURL"],
	)
}

func (*Chinese) Text_ValueStatistics(values map[string]interface{}) string {
	return fmt.Sprintf(`
	共%d个导出变量和%d个导出常量。

	<img src="%s"></image>
	<img src="%s"></image>

	共%d个导出函数和%d个导出显式方法。
	平均说来，每个这样的函数或方法拥有%.2f个参数和%.2f个输出结果。
	这些函数和方法中的%d个（占%d%%）的最后一个输出结果的类型为<code>error</code>。

	<img src="%s"></image>
	<img src="%s"></image>
`,
		values["exportedVariables"],
		values["exportedConstants"],

		values["exportedvariablesByTypekindsChartURL"],
		values["exportedconstantsByTypekindsChartURL"],

		values["exportedFunctions"],
		values["exportedMethods"],
		values["averageParameterCountPerExportedFunction"],
		values["averageResultCountPerExportedFunction"],
		values["exportedFunctionWithLastErrorResult"],
		values["exportedFunctionWithLastErrorResultPercentage"],

		values["exportedfunctionsByParametersChartURL"],
		values["exportedfunctionsByResultsChartURL"],
	)
}

func (*Chinese) Text_Othertatistics(values map[string]interface{}) string {
	return fmt.Sprintf(`
	输出标识符的平均长度为%.2f。

	<img src="%s"></image>
`,
		values["averageIdentiferLength"],

		values["exportedidentifiersByLengthsChartURL"],
	)
}

///////////////////////////////////////////////////////////////////
// footer
///////////////////////////////////////////////////////////////////

func (*Chinese) Text_GeneratedPageFooter(goldVersion, qrCodeLink, goOS, goArch string) string {
	var qrImg string
	if qrCodeLink != "" {
		qrImg = fmt.Sprintf(`<img src="%s">`, qrCodeLink)
	}
	return fmt.Sprintf(`<table><tr><td>%s</td>
<td>本页面由 <a href="https://go101.org/article/tool-gold.html"><b>Gold</b></a> <i>%s</i> 生成。（GOOS=%s GOARCH=%s）。
<b>Gold</b> 是由<a href="https://gfw.tapirgames.com">老貘</a>创建的一个 <a href="https://go101.org">Go 101</a>项目。
欢迎在 <a href="https://github.com/go101/gold">Gold 项目</a>中提交 PR 和 bug 报告。
请关注 “Go 101” 微信公众号（扫描左边的二维码）以获取 <b>Gold</b> 的最新消息以及各种 Go 细节和事实。</td></tr></table>`,
		qrImg,
		goldVersion,
		goOS,
		goArch,
	)
}
