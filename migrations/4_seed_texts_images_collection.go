package migrations

import (
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		create_table_texte(db)
		create_table_bilder(db)

		seed_texte(db)
		return nil
	}, func(db dbx.Builder) error { // revert op
		dao := daos.New(db)
		err := dao.DeleteTable("Texte")
		if err != nil {
			fmt.Println(err)
		}

		err = dao.DeleteTable("Bilder")
		if err != nil {
			fmt.Println(err)
		}

		return nil
	})
}

func create_table_texte(db dbx.Builder) error {
	dao := daos.New(db)
	collection := &models.Collection{
		Name:       "Texte",
		Type:       models.CollectionTypeBase,
		System:     true,
		ListRule:   types.Pointer(""),
		ViewRule:   types.Pointer(""),
		CreateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
		UpdateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
		DeleteRule: nil,
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:        "Titel",
				Type:        schema.FieldTypeText,
				Required:    true,
				Presentable: true,
			},
			&schema.SchemaField{
				Name:        "Stichworte",
				Type:        schema.FieldTypeText,
				Required:    false,
				Presentable: true,
			},
			&schema.SchemaField{
				Name:     "Text",
				Type:     schema.FieldTypeEditor,
				Required: false,
			},
		),
	}

	if err := dao.SaveCollection(collection); err != nil {
		return err
	}

	return nil
}

func create_table_bilder(db dbx.Builder) error {
	dao := daos.New(db)
	collection := &models.Collection{
		Name:       "Bilder",
		Type:       models.CollectionTypeBase,
		System:     true,
		ListRule:   types.Pointer(""),
		ViewRule:   types.Pointer(""),
		CreateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
		UpdateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:        "Titel",
				Type:        schema.FieldTypeText,
				Required:    true,
				Presentable: true,
			},
			&schema.SchemaField{
				Name:     "Beschreibung",
				Type:     schema.FieldTypeEditor,
				Required: false,
			},
			&schema.SchemaField{
				Name:        "Vorschau",
				Type:        schema.FieldTypeFile,
				Required:    false,
				Presentable: false,
				Options: &schema.FileOptions{
					MaxSelect: 1,
					MaxSize:   524288000,
					MimeTypes: []string{
						"image/png",
						"image/vnd.mozilla.apng",
						"image/jpeg",
						"image/jp2",
						"image/jpx",
						"image/jpm",
						"image/gif",
						"image/jxs",
						"image/jxl",
						"image/x-xpixmap",
						"image/vnd.adobe.photoshop",
						"image/webp",
						"image/tiff",
						"image/bmp",
						"image/x-icon",
						"image/vnd.djvu",
						"image/bpg",
						"image/vnd.dwg",
						"image/x-icns",
						"image/heic",
						"image/heic-sequence",
						"image/heif",
						"image/heif-sequence",
						"image/vnd.radiance",
						"image/x-xcf",
						"image/x-gimp-pat",
						"image/x-gimp-gbr",
						"image/avif",
						"image/jxr",
						"image/svg+xml",
					},
					Thumbs: []string{
						"100x0",
						"250x0",
						"500x0",
						"0x250",
						"0x500",
					},
					Protected: false,
				},
			},
			&schema.SchemaField{
				Name:        "Bilder",
				Type:        schema.FieldTypeFile,
				Required:    false,
				Presentable: false,
				Options: &schema.FileOptions{
					MaxSelect: 1000,
					MaxSize:   524288000,
					MimeTypes: []string{
						"image/png",
						"image/vnd.mozilla.apng",
						"image/jpeg",
						"image/jp2",
						"image/jpx",
						"image/jpm",
						"image/gif",
						"image/jxs",
						"image/jxl",
						"image/x-xpixmap",
						"image/vnd.adobe.photoshop",
						"image/webp",
						"image/tiff",
						"image/bmp",
						"image/x-icon",
						"image/vnd.djvu",
						"image/bpg",
						"image/vnd.dwg",
						"image/x-icns",
						"image/heic",
						"image/heic-sequence",
						"image/heif",
						"image/heif-sequence",
						"image/vnd.radiance",
						"image/x-xcf",
						"image/x-gimp-pat",
						"image/x-gimp-gbr",
						"image/avif",
						"image/jxr",
						"image/svg+xml",
					},
					Thumbs: []string{
						"100x0",
						"250x0",
						"500x0",
						"0x250",
						"0x500",
					},
					Protected: false,
				},
			},
		),
	}

	if err := dao.SaveCollection(collection); err != nil {
		return err
	}

	return nil
}

func seed_texte(db dbx.Builder) error {
	dao := daos.New(db)
	collection, err := dao.FindCollectionByNameOrId("Texte")
	if err != nil {
		fmt.Println(err)
		return err
	}

	model_start := models.NewRecord(collection)
	model_start.Set("Titel", "Start")
	model_start.Set("Text", "<p>Ziel der Musenalm ist die&nbsp;bibliographische Erfassung eines Jahrhunderts deutscher Almanache und Taschenb&uuml;cher;<strong>&nbsp;</strong>das Projekt ist im Aufbau und wird kontinuierlich weitergef&uuml;hrt.</p><p>Verzeichnet werden:</p><ul><li><strong>Reihen </strong>und<strong> B&auml;nde</strong> bekannter Almanache und einzelne Druckauflagen mit ausf&uuml;hrlichen bibliographischen Angaben und kurzer systematisierter&nbsp;<strong>Darstellung ihres strukturellen Aufbaus </strong>&nbsp;(Paginierung, Anordnung der Druckteile, Graphiken und Beilagen),<strong><br></strong></li><li><strong>Beitr&auml;ge literarisch oder musisch ausgerichteter Almanache&nbsp;</strong>einzeln, nach Autor, &Uuml;berschrift und Incipit,<strong> </strong>inklusive<strong> Digitalisate </strong>graphischer und musischer Beitr&auml;ge,</li><li>Beitr&auml;ge vorwiegend&nbsp;<strong>nicht literarischer Almanache</strong>&nbsp;in der Regel durch Wiedergabe des&nbsp;<strong>Inhaltsverzeichnisses.</strong></li></ul><p>Die Bibliographie ist zug&auml;nglich mit umfangreichen Suchfunktionen &uuml;ber:</p><ul><li><strong>Reihentitel der Almanache,</strong></li><li><strong>Abbildungen (Graphiken und Musikbeilagen),</strong></li><li>Personennamen von Herausgebern und Beitr&auml;gern einerseits &uuml;ber normierte<strong> Realnamen </strong>und andererseits &uuml;ber die im Druck erscheinenden Schreibweisen der Personen (auch Pseudonyme)<strong> </strong>als<strong> Autornamen,</strong></li><li><strong>Einzeltitel und Incipit </strong>(w&ouml;rtliche Textanf&auml;nge) von Beitr&auml;gen.</li></ul><p>Die Musenalm ist ein Projekt der Theodor Springmann Stiftung in Heidelberg.</p>")
	if err := dao.SaveRecord(model_start); err != nil {
		fmt.Println(err)
	}

	model_einf := models.NewRecord(collection)
	model_einf.Set("Titel", "Einführung")
	model_einf.Set("Text", "<h2 class=\"wp-block-heading\">Vorbemerkung</h2>\r\n<p>Dies ist eine Bibliographie der deutschen Almanache und Taschenb&uuml;cher, die neben der Erfassung der Reihen und ihrer Jahrg&auml;nge die Inhalte selbst erkennbar macht. In der Regel werden folgende Merkmale erfa&szlig;t und sind in verschiedenen Suchabfragen und Listen abrufbar:</p>\r\n<ul>\r\n<li>Reihen- und Einzeltitel des Druckwerks sowie Strukturdarstellung des autopsierten Einzelbandes.</li>\r\n<li>Namen der Herausgeber und Verfasser, gegebenfalls zus&auml;tzlich Schreibvarianten oder Pseudonyme.</li>\r\n<li>Literarische Beitr&auml;ge mit Titel und Incipit.</li>\r\n<li>Nichtliterarische Beitr&auml;ge (Illustrationen, Musikbeilagen und andere Zutaten) werden als Vollbild gegeben.</li>\r\n</ul>\r\n<p>Erfa&szlig;t werden in erster Linie die literarischen Taschenb&uuml;cher, die von 1770 bis etwa 1870 erschienen sind. Angesichts der raschen modischen Entwicklung des Almanachwesens, das sich schnell auf viele und auch entlegene Themengebiete ausdehnte, ist eine klare Abgrenzung der literarischen zu anders ausgerichteten Erscheinungen schwierig und wird von uns nicht angestrebt. Vielmehr sind wir bem&uuml;ht, das ganze Spektrum des Almanachwesens sichtbar zu machen, und wir nehmen ebenfalls, wenn auch zun&auml;chst nur ausgew&auml;hlt und nicht vollz&auml;hlig, unliterarische Taschenbuchreihen auf, die wir zumeist allerdings nicht in inhaltlicher Aufgliederung, sondern nur unter dem Titel bibliographieren, unter Beif&uuml;gung einer kurzen allgemeinen Beschreibung. Graphische Darstellungen in solchen Reihen sollen jedoch ebenfalls bildlich aufgenommen werden, sofern sie nicht als vergleichsweise unbedeutend erscheinen.</p>\r\n<p>Ausgegegrenzt bleiben die eigentlichen Land- und Volkskalender, die vorwiegend im Quart-Format, im Verlauf des 19. Jahrhunderts zunehmend aber auch im Oktav-Format erschienen sind.</p>\r\n<p>Das Vorhaben wird von der <a href=\"https://musenalm.de/kontakt.html\">THEODOR SPRINGMANN STIFTUNG</a> betrieben und greift zun&auml;chst auf deren umfangreichen Sammlungsbestand zur&uuml;ck, im weiteren Verlauf werden wir auf die Einbeziehung anderer Bibliotheken nicht verzichten k&ouml;nnen; vielfach wird es auch n&ouml;tig sein, M&auml;ngel und Fehlstellen einzelner vorhandener St&uuml;cke zu erg&auml;nzen.</p>\r\n<p>Wir bitten um Anregungen und Korrekturen. Auch Hilfen durch methodische bibliographische Aufnahmen sind hochwillkommen und tragen zur Verk&uuml;rzung des langwierigen Verfahrens bei. Hierzu k&ouml;nnen entsprechende Formulare bei uns angefordert werden.</p>\r\n<p>Das Inhaltsverzeichnis der deutschen Almanache wird erarbeitet von Wolfgang Binnig und Martin Sietzen und herausgegeben von Adrian Braunbehrens.</p>\r\n<h2 class=\"wp-block-heading\">Einleitung in das<br>Inhaltsverzeichnis deutscher Almanache</h2>\r\n<p>Seit Kalender geschrieben und gedruckt wurden, finden wir sie verquickt mit anderen Momenten der Jahreszeitlichkeit. Hierzu z&auml;hlen astronomische und astrologische Angaben, die Nennung guter und b&ouml;ser Tage, praktische Regeln zu den Jahreszeiten und ihrer Witterung und nicht zuletzt Texte zu musischem und geselligem Zeitvertreib. Dies f&uuml;hrte zur Ausbildung besonderer Typen, die einzelne dieser Momente ausf&uuml;hrlicher vorstellten. Zu den eigenartigsten und reizvollsten z&auml;hlen die poetischen Musenalmanache und literarisch unterhaltenden Taschenb&uuml;cher. Ihre Epoche begann in Deutschland &ndash; franz&ouml;sischen Vorbildern folgend &ndash; um 1770 und endete gegen 1848. Sie wurden zur wohl h&uuml;bschesten und zugleich langlebigsten Modeerscheinung auf dem deutschen Buchmarkt.</p>\r\n<p>In Paris erschien 1765 der ALMANACH DES MUSES OU CHOIX DE PO&Eacute;SIES FUGITIVES, eine Bl&uuml;tenlese neuerer Dichtung, dessen Reihe sich in j&auml;hrlicher Ausgabe bis ins 19. Jahrhundert fortsetzte. Er fand alsbald auch in Deutschland Beachtung und zum Jahr 1770 erschien, herausgegeben von Heinrich Christian Boie und verlegt von J. C. Dieterich der G&Ouml;TTINGER MUSENALMANACH. Klopstock z&auml;hlte neben vielen anderen zu seinen Beitr&auml;gern. Der 1772 gegr&uuml;ndete G&ouml;ttinger Dichterbund fand in ihm ein Organ. Johann Heinrich Voss &uuml;bernahm 1774 f&uuml;r kurze Zeit die Redaktion, begann indes 1776 in Hamburg mit einem eigenen Musenalmanach, der nun &ndash; in recht &auml;hnlichem Erscheinungsbild &ndash; neben dem G&ouml;ttinger bestand. Von diesen Keimzellen aus entwickelte sich in Deutschland ein literarisches Almanach- und Taschenbuchwesen, das in vief&auml;ltiger Unterschiedenheit eine kaum &uuml;bersehbare, nach Tausenden zu z&auml;hlende, oft kurzlebige, teils auch in langen Jahresreihen sich fortsetzende F&uuml;lle hervorbrachte. Die Epoche des literarischen Taschenbuches war zugleich eine Hochzeit dichterischer Entfaltung in Deutschland, in den H&ouml;hen neuer Findung ebenso wie in den Senken der Trivialit&auml;t; und es war eine Periode umfassender sozialer Ver&auml;nderungen. Das Taschenbuch als gesellschaftliche Modeerscheinung und seine Wandlungen stehen in engem Bezug zu diesen Entwicklungen.- Eine umfassende Geschichte des Taschenbuches unter ausf&uuml;hrlicher Ber&uuml;cksichtigung dieser Bez&uuml;ge ist noch nicht geschrieben worden.</p>\r\n<p>Die Mannigfaltigkeit des periodischen Taschenbuches zeigt so viele Facetten, da&szlig; es schwer h&auml;lt, eine best&auml;ndige Gattung auszumachen. Beschreiben lassen sich wiederkehrende Einzelmerkmale, die in unterschiedlicher H&auml;ufung, kaum aber in ihrer Gesamtheit beim jeweiligen Exemplar vorkommen. Unsicher ist schon die Verwendung der Ausdr&uuml;cke ALMANACH und TASCHENBUCH oder auch KALENDER; sie &uuml;berschneiden sich gro&szlig;enteils, ohne sich ganz zu decken. Es k&ouml;nnte sich empfehlen, im Taschenbuch den Oberbegriff zu sehen, wenn nicht heutzutage unter dem Taschenbuch eine ganz andere Produktform des Buches bezeichnet w&auml;re.</p>\r\n<p>Das Wort ALMANACH (arabischen Ursprungs) ist eine Bezeichnung f&uuml;r Kalender, und mit dem Kalender hat das hier dargestellte Taschenbuch die angelegte J&auml;hrlichkeit gemein, auch wenn so manche Erscheinung &uuml;ber den ersten Jahrgang nicht hinauskommt. Oftmals, aber keineswegs immer und immer seltener werdend, ist ein Kalender dem Textteil vorgebunden. Regional erhobene Kalender-Stempelsteuern konnten hier prohibitiv wirken. Einige besonders erfolgreiche Almanache erfuhren noch Jahre nach dem Erstdruck Folgeauflagen, in denen dann der &uuml;berfl&uuml;ssig gewordene Kalender, nicht jedoch die urspr&uuml;ngliche Jahresdatierung, entfallen konnte. &ndash; Seiner Entwicklungsgeschichte nach ist das Taschenbuch durchaus vom Kalender herzuleiten, aber es emanzipiert sich gleichsam von diesem und l&auml;&szlig;t seine Herkunft vergessen. Was bleibt ist die Taschenhandlichkeit des Formates: Sedez oder Duodez, seltener Octav (aber auch hierzu in der Sp&auml;tzeit die seltene Ausnahme des Gro&szlig;octav). Und es scheint, da&szlig; die Almanache, Kalendern gleich, meist keinen Ruheplatz in den B&uuml;cherschr&auml;nken gefunden haben, sondern zur Hand genommen und vernutzt wurden; die bis heute erhalten gebliebenen Exemplare sind nicht selten ramponiert, zum Schaden f&uuml;r den zierlich gestalteten Einband.</p>\r\n<p>Welche Art von Texten f&uuml;llte die Almanache und Taschenb&uuml;cher? Anfangs war es Lyrik, sehr bald aber, als die Mode grassierte: quodlibet, alles was beliebt; unterhalten sollte es, in Spa&szlig; oder Ernst. Nur selten mischt Belehrendes sich ein, im Unterschied zum gr&ouml;&szlig;er formatierten aber sehr viel schmaleren Land- oder Volkskalender. Sieht man in das Register der vorz&uuml;glichen <em>Geschichte der deutschen Taschenb&uuml;cher und Almanache aus der klassisch-romantischen Zeit</em> von LANCKORONSKA und R&Uuml;MANN, so findet man schon in den Titeln die Hinweise auf jede nur denkbare Art von Adressaten und zugeh&ouml;rigen Inhalten: Wanderer, Reiter, Bienenfreunde, K&uuml;nstler, Scheidek&uuml;nstler und Apotheker, Liebende, Tollh&auml;usler, Ketzer, &Auml;rzte und Nicht&auml;rzte, Charadenfreunde, Kaufleute, Lottospieler u.v.a.m.. Vor allem aber wird die Weiblichkeit angesprochen, seien es Frauenzimmer oder Damen, Dienstm&auml;dchen, das Sch&ouml;ne Geschlecht, Kammerjungfern, Grabennymphen, Edle Weiber und M&auml;dchen. Selbst wenn es der Titel nicht verr&auml;t, ist &ouml;fter an die Leserin gedacht als an den Herrn, sie hatte wohl mehr gesellige Mu&szlig;e, und sie war der gemeinte Empf&auml;nger des h&uuml;bschen kleinen Geschenks. Denn zum Schenken war er bestimmt und dazu f&uuml;gte sich der Erscheinungstermin zur Michaelismesse, rechtzeitig zu Weihnachten und Neujahr.</p>\r\n<p>Schwerpunkt der bibliographischen Erfassung und inhaltlichen Erschlie&szlig;ung sind zun&auml;chst die literarischen Almanache &ndash; ungeachtet ihres Niveaus. Sie sind Versammlungsort nicht nur der Gro&szlig;en, sondern vorz&uuml;glich derjenigen Dichter und Prosaisten, deren Schriften heute &ndash; zu Recht oder zu Unrecht&ndash; vergessen sind, die aber aus manchen Gr&uuml;nden gelegentlich doch in den Blick des Interesses r&uuml;cken. Das Verzeichnis soll sie, die bislang nur unter Schwierigkeiten aufzufinden waren, zug&auml;nglich machen. Besonders wichtig, weil eine Wahrnehmungsl&uuml;cke f&uuml;llend, erschien uns daneben die Registrierung der Zeichner und Stecher, deren Graphiken wir als Vollbild wiedergeben wollen. Da&szlig; gerade in diesem Bereich die vorliegenden Exemplare oft unvollst&auml;ndig sind, f&uuml;hrt gelegentlich zu Fehlstellen in unserer Darstellung (die aber auf Dauer geschlossen werden); es unterstreicht zugleich die Notwendigkeit des gesetzten Ziels. Indes werden nicht nur die Vorlagen M&auml;ngel aufweisen, auch in der Bearbeitung werden unvermeidbar Fehler entstehen. Wir bitten aufmerksame Benutzer, uns hier&uuml;ber zu informieren und dadurch zur Besserung zu verhelfen.</p>\r\n<p>Auf l&auml;ngere Sicht sollen alle periodisch angelegten Almanache und Taschenb&uuml;cher des 18. und 19. Jahrhunderts aufgenommen werden, um das gesamte Spektrum dieser Publikationsart sichtbar zu machen. Im nicht-literarischen Bereich werden wir uns jedoch zumeist beschr&auml;nken auf die bibliographische Registrierung und eine kurze Beschreibung der Einzelb&auml;nde und wir werden hierbei auf die ausf&uuml;hrliche Inhalts&uuml;bersicht verzichten und uns mit der Wiedergabe eines Inhaltsverzeichnisses begn&uuml;gen.</p>\r\n<p>Grunds&auml;tzlich ist Voraussetzung unserer bibliographischen Erfassung die Autopsie des Einzelemplares. Dies sch&uuml;tzt indes nicht immer vor Verwirrung: Variante Doppeldrucke (etwa bei unbezeichnetet Folgeauflagen oder nach Zensureingriffen), fehlende Bl&auml;tter und andere Fehlerquellen sind nicht in jedem Fall wahrnehmbar. Auf alles auff&auml;llig Sonderliche wird anmerkend hingewiesen. Um uns m&ouml;glicher Vollst&auml;ndigkeit anzun&auml;hern, behalten wir uns vor, im Einzelfall auch ohne Autopsie nach bibliographischen Vorgaben aufzunehmen; wir werden dies jedoch immer unter Nennung der Quelle ausdr&uuml;cklich anmerken.</p>\r\n<p>Adrian Braunbehrens</p>")
	if err := dao.SaveRecord(model_einf); err != nil {
		fmt.Println(err)
	}

	model_b_reihe := models.NewRecord(collection)
	model_b_reihe.Set("Titel", "Begriffserklärung: Reihe")
	model_b_reihe.Set("Text", "<p>Die Reihe ist der Oberbegriff f&uuml;r eine Anzahl zusammengeh&ouml;riger Almanache, die durch Konzept, Titel, Herausgeber und fortlaufendes Erscheinen miteinander verbunden sind. H&auml;ufig sind jedoch Wechsel der Herausgeber, Varianz im Titel, unregelm&auml;&szlig;iges Erscheinen zu beobachten, so da&szlig;, um Einheitlichkeit und Suchbarkeit zu erm&ouml;glichen, von den konkreten Einzeltiteln der Reihentitel abstrahiert werden mu&szlig;.</p>")
	if err := dao.SaveRecord(model_b_reihe); err != nil {
		fmt.Println(err)
	}

	model_b_band := models.NewRecord(collection)
	model_b_band.Set("Titel", "Begriffserklärung: Band")
	model_b_band.Set("Text", "<p>Der einzelne Almanach (das einzelne Buch) ist ein bestimmter Band oder Jahrgang einer Reihe.<br>Angegeben werden Titel und Reihentitel des Bandes, angezeigter Name und Realname des Herausgebers, Ort und Jahr, Angaben zum Aufbau, Anmerkungen und die interne Nummer des Almanachs.</p>")
	if err := dao.SaveRecord(model_b_band); err != nil {
		fmt.Println(err)
	}

	model_b_inhalt := models.NewRecord(collection)
	model_b_inhalt.Set("Titel", "Begriffserklärung: Inhalt")
	model_b_inhalt.Set("Text", "<p>Der Inhalt eines Almanachs ist die Zusammenstellung der einzelnen Beitr&auml;ge, die in einem bestimmten Band erschienen sind. Erfasst werden Beitr&auml;ge nach zahlreichen Kriterien, wie Name  des Autors, Titel und Incipit, Art des Beitrags, Angaben zum Ort im Almanach, Anmerkungen.</p>")
	if err := dao.SaveRecord(model_b_inhalt); err != nil {
		fmt.Println(err)
	}

	model_dok_datenfelder := models.NewRecord(collection)
	model_dok_datenfelder.Set("Titel", "Dokumentation: Datenbankfelder")
	model_dok_datenfelder.Set("Text", "<table>\r\n<tbody>\r\n<tr>\r\n<th><strong>Feldname</strong></th>\r\n<th><strong>Datenbank</strong></th>\r\n<th><strong>Ergebnisseite</strong></th>\r\n<th><strong>Beschreibung</strong></th>\r\n</tr>\r\n<tr>\r\n<th><strong>Almanache</strong></th>\r\n</tr>\r\n<tr>\r\n<td><strong>Titel</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Titel des Almanachs ohne &Auml;nderungen, Auslassungen oder K&uuml;rzungen. Schreibweise: wie im Almanach</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Reihentitel</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Der Reihentitel wird pragmatisch vergeben, er dient dazu, zusammengeh&ouml;rige B&auml;nde trotz &Auml;nderungen des Titels etc. unter einem einheitlichen Namen zu erfassen. Der Reihentitel, auch Kurztitel genannt, dient als Suchtitel.</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Herausgeber</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Name des Herausgebers, wie im Almanach zu finden</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Realname</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Der gedeutete Name in &uuml;berlieferter Schreibweise</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Ort</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Verlagsort(e) wie im Almanach angegeben.</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Jahr</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Jahr, auf das sich der Almanach im Titel bezieht.</td>\r\n</tr>\r\n<tr>\r\n<td><strong>AlmanachNr</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Eindeutige Referenznummer des Almanachs</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Struktur*</strong></td>\r\n<td>nein</td>\r\n<td>ja</td>\r\n<td>Reihenfolge der tats&auml;chlich in diesem Band vorliegenden Inhaltsobjekte, wobei diese nur nach ihrer Kategorie, nicht nach den Details, aufgelistet werden.</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Norm*</strong></td>\r\n<td>nein</td>\r\n<td>ja</td>\r\n<td>Reihenfolge der tats&auml;chlichen oder vermutlich beabsichtigten Inhaltsobjekte; Aufbau des Almanachs</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Anmerkungen</strong></td>\r\n<td>nein</td>\r\n<td>ja</td>\r\n<td>Anmerkungen zum Band. Im ersten Band einer Reihe finden sich Angaben zu Herausgeber, Ort, Verlag und Erscheinungsfolge.</td>\r\n</tr>\r\n<tr>\r\n<th><strong>Inhalte</strong></th>\r\n</tr>\r\n<tr>\r\n<td><strong>Autor</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Der angezeigte Name des Autors (auch Pseudonyme und K&uuml;rzel oder &raquo;unbezeichnet&laquo;), wie im Almanach zu finden</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Realname</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Der gedeutete Name in &uuml;berlieferter Schreibweise</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Titel</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Titel des Objekts</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Incipit</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Textanfang (ca 2 Zeilen)</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Objekt</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Um was handelt es sich? Vgl. unten die Tabelle&nbsp;<strong>Objekte</strong></td>\r\n</tr>\r\n<tr>\r\n<td><strong>Abbildung</strong></td>\r\n<td>nein</td>\r\n<td>ja</td>\r\n<td>Ja/Nein f&uuml;r Foto des Objekts vorhanden/nicht vorhanden.</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Paginierung</strong></td>\r\n<td>nein</td>\r\n<td>ja</td>\r\n<td>r&ouml;mische/arabische Paginierung</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Seite</strong></td>\r\n<td>nein</td>\r\n<td>ja</td>\r\n<td>Seitennummer nach arabischer oder r&ouml;mischer Paginierung</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Objektz&auml;hler</strong></td>\r\n<td>nein</td>\r\n<td>ja</td>\r\n<td>Unabh&auml;ngig von Art oder vorhandener Paginierung wird jedem Inhalt seine relative Position zugewiesen.</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Inhaltsnummer</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Eindeutige Datensatznummer</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Anmerkungen</strong></td>\r\n<td>ja</td>\r\n<td>ja</td>\r\n<td>Anmerkungen zum jeweiligen Objekt</td>\r\n</tr>\r\n</tbody>\r\n</table>")
	if err := dao.SaveRecord(model_dok_datenfelder); err != nil {
		fmt.Println(err)
	}

	model_dok_sym := models.NewRecord(collection)
	model_dok_sym.Set("Titel", "Dokumentation: Symbole/Abkürzungen")
	model_dok_sym.Set("Text", "<table>\r\n<tbody>\r\n<tr>\r\n<td><strong>Anm.</strong></td>\r\n<td>Anmerkung</td>\r\n</tr>\r\n<tr>\r\n<td><strong>ar, ar1, ar2</strong></td>\r\n<td>arabische Paginierung (ggf mehrere)</td>\r\n</tr>\r\n<tr>\r\n<td><strong>B; BB</strong></td>\r\n<td>Blatt; Bl&auml;tter</td>\r\n</tr>\r\n<tr>\r\n<td><strong>C</strong></td>\r\n<td>Corrigenda</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Diagr</strong></td>\r\n<td>Diagramm</td>\r\n</tr>\r\n<tr>\r\n<td><strong>G; GG</strong></td>\r\n<td>Graphik; Graphiken</td>\r\n</tr>\r\n<tr>\r\n<td><strong>UG r, v</strong></td>\r\n<td>Umschlaggraphik recto, verso</td>\r\n</tr>\r\n<tr>\r\n<td><strong>TG r, v</strong></td>\r\n<td>Titelgraphik, Titelportrait etc</td>\r\n</tr>\r\n<tr>\r\n<td><strong>gA</strong></td>\r\n<td>graphische Anleitung</td>\r\n</tr>\r\n<tr>\r\n<td><strong>gTzA</strong></td>\r\n<td>graphische Tanzanleitung</td>\r\n</tr>\r\n<tr>\r\n<td><strong>G-Verz</strong></td>\r\n<td>Verzeichnis der Kupfer u. &auml;.</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Hrsg</strong></td>\r\n<td>Herausgeber</td>\r\n</tr>\r\n<tr>\r\n<td><strong>I-Verz</strong></td>\r\n<td>Inhaltsverzeichnis</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Kal</strong></td>\r\n<td>Kalendarium</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Kr</strong></td>\r\n<td>Karte</td>\r\n</tr>\r\n<tr>\r\n<td><strong>MusB; MusBB</strong></td>\r\n<td>Musikbeigabe; Musikbeigaben</td>\r\n</tr>\r\n<tr>\r\n<td><strong>r&ouml;m, r&ouml;m1, r&ouml;m2</strong></td>\r\n<td>r&ouml;mische Paginierung (ggf mehrere)</td>\r\n</tr>\r\n<tr>\r\n<td><strong>S; SS</strong></td>\r\n<td>Seite; Seiten</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Sp</strong></td>\r\n<td>Spiegel</td>\r\n</tr>\r\n<tr>\r\n<td><strong>T</strong></td>\r\n<td>Titel</td>\r\n</tr>\r\n<tr>\r\n<td><strong>gT</strong></td>\r\n<td>graphischer Titel</td>\r\n</tr>\r\n<tr>\r\n<td><strong>vT</strong></td>\r\n<td>Vortitel</td>\r\n</tr>\r\n<tr>\r\n<td><strong>nT</strong></td>\r\n<td>Nachtitel</td>\r\n</tr>\r\n<tr>\r\n<td><strong>zT</strong></td>\r\n<td>Zwischentitel</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Tab</strong></td>\r\n<td>Tabelle</td>\r\n</tr>\r\n<tr>\r\n<td><strong>VB</strong></td>\r\n<td>Vorsatzblatt</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Vf</strong></td>\r\n<td>Verfasser</td>\r\n</tr>\r\n<tr>\r\n<td><strong>VrlgM</strong></td>\r\n<td>Verlagsmitteilung</td>\r\n</tr>\r\n<tr>\r\n<td><strong>Vrwrt</strong></td>\r\n<td>Vorwort</td>\r\n</tr>\r\n<tr>\r\n<td><strong>#</strong></td>\r\n<td>Hinweis auf weitere Informationen in der Anmerkung.</td>\r\n</tr>\r\n<tr>\r\n<td><strong>&sect;&sect;</strong></td>\r\n<td>Hinweis auf M&auml;ngel im Almanach (Besch&auml;digungen, fehlende Graphiken od. Beitr&auml;ge, unvollst&auml;ndige Sammlungen etc) in der Anmerkung</td>\r\n</tr>\r\n<tr>\r\n<td><strong>+++</strong></td>\r\n<td>Inhalte aus mehreren Almanachen interpoliert</td>\r\n</tr>\r\n<tr>\r\n<td><strong>$</strong></td>\r\n<td>vermutlich</td>\r\n</tr>\r\n</tbody>\r\n</table>")
	if err := dao.SaveRecord(model_dok_sym); err != nil {
		fmt.Println(err)
	}

	model_lit_siglen := models.NewRecord(collection)
	model_lit_siglen.Set("Titel", "Literatur: Siglen")
	model_lit_siglen.Set("Text", "<table>\r\n<tbody>\r\n<tr>\r\n<th>Sigle</th>\r\n<th>Titel</th>\r\n</tr>\r\n<tr>\r\n<td>ADB</td>\r\n<td>Historische Commission bei der K&ouml;niglichen Akademie der Wissenschaften (Hg.):&nbsp;<em>Allgemeine Deutsche Biographie.</em>&nbsp;55 Bde., Leipzig 1875-1910.</td>\r\n</tr>\r\n<tr>\r\n<td>Brockh (13.)</td>\r\n<td><em>Brockhaus&lsquo; Conversations-Lexikon. Allgemeine deutsche Real-Encyklop&auml;die.</em>&nbsp;13. vollst. umgearb. Aufl., Leipzig 1883-1888.</td>\r\n</tr>\r\n<tr>\r\n<td>D</td>\r\n<td>Diesch, Carl:&nbsp;<em>Bibliographie der germanistischen Zeitschriften.</em>&nbsp;Leipzig 1927. (= Heuser, Frederick W. J. [Hg.]:&nbsp;<em>Bibliographical Publications.</em>&nbsp;Bd. 1: Diesch, Carl:&nbsp;<em>Bibliography of Germanic Periodicals.</em>&nbsp;Leipzig 1927.)</td>\r\n</tr>\r\n<tr>\r\n<td>DBI</td>\r\n<td>Gorzny, Willy (Hg.):&nbsp;<em>Deutscher Biographischer Index</em>.Bearbeitet von Hans-Albrecht Koch, Uta Koch und Angelika Koller, M&uuml;nchen 1986. [Register zu: Fabian, Bernhard (Hg.):&nbsp;<em>Deutsches Biographisches Archiv.</em>&nbsp;M&uuml;nchen 1960-1999.]</td>\r\n</tr>\r\n<tr>\r\n<td>E</td>\r\n<td>Eymer, Wilfried:&nbsp;<em>Eymers Pseudonymen Lexikon. Realnamen und Pseudonyme in der deutschen Literatur</em>. Bonn 1997.</td>\r\n</tr>\r\n<tr>\r\n<td>G</td>\r\n<td>Goldschmidt, Arthur:&nbsp;<em>Goethe im Almanach.</em>&nbsp;Leipzig 1932.</td>\r\n</tr>\r\n<tr>\r\n<td>G-C</td>\r\n<td>Grand-Carteret, John:&nbsp;<em>Les Almanachs Francais, Bibliographie &ndash; Iconographie 1600-1895</em>. Paris 1896.</td>\r\n</tr>\r\n<tr>\r\n<td>Goed</td>\r\n<td>Goedeke, Carl u. a.:&nbsp;<em>Grundri&szlig; zur Geschichte der deutschen Dichtung.</em>&nbsp;13 Bde., 2. Aufl., Dresden 1884 ff. (Bd. IV/1-4 in der dritten neubearbeiteten Aufl., Dresden 1896-1913. Neudruck Berlin 1955; Bd. IV/5 in der ersten Aufl. D&uuml;sseldorf 1957-1960; au&szlig;erdem die &bdquo;Neue Folge&ldquo;, N. F. Bd. I, Berlin 1955ff.)</td>\r\n</tr>\r\n<tr>\r\n<td>Goed (H)</td>\r\n<td>Hirschberg, Leopold:&nbsp;<em>Der Taschengoedeke. Bibliographie deutscher Erstausgaben.</em>&nbsp;M&uuml;nchen 1970.</td>\r\n</tr>\r\n<tr>\r\n<td>H/B</td>\r\n<td>Holzmann, Michael; Bohatta, Hanns:&nbsp;<em>Deutsches Pseudonymen-Lexikon. Aus den Quellen bearbeitet von Michael Holzmann und Hanns Bohatta</em>. Hildesheim, New York. 1970.</td>\r\n</tr>\r\n<tr>\r\n<td>K</td>\r\n<td>K&ouml;hring, Hans (Bearb. Hrsg.):&nbsp;<em>Bibliographie der Almanache, Kalender und Taschenb&uuml;cher f&uuml;r die Zeit von ca. 1750-1860. Hamburg 1929.</em>&nbsp;Neudruck Bad Karlshafen 1987.</td>\r\n</tr>\r\n<tr>\r\n<td>L/R</td>\r\n<td>Lanckoronska, Maria; R&uuml;mann, Arthur:&nbsp;<em>Geschichte der deutschen Taschenb&uuml;cher und Almanache aus der klassisch-romantischen Zeit.</em>&nbsp;M&uuml;nchen 1957. Neudruck Osnabr&uuml;ck 1985.</td>\r\n</tr>\r\n<tr>\r\n<td>NBG</td>\r\n<td>Hoefer, Jean Chr&eacute;tien Ferdinand (Hg.):&nbsp;<em>Nouvelle biographie g&eacute;n&eacute;rale.</em>&nbsp;46 Bde. Paris 1851-1878.</td>\r\n</tr>\r\n<tr>\r\n<td>P</td>\r\n<td>Ziegler, Konrat; Sontheimer, Walther (Hg.):&nbsp;<em>Der kleine Pauly. Lexikon der Antike</em>. 5 Bde., M&uuml;nchen 1979.</td>\r\n</tr>\r\n<tr>\r\n<td>R&uuml;</td>\r\n<td>R&uuml;mann, Arthur:&nbsp;<em>Die illustrierten deutschen B&uuml;cher des 18. Jahrhunderts.</em>&nbsp;Stuttgart 1927.</td>\r\n</tr>\r\n<tr>\r\n<td>Re</td>\r\n<td>Redlich, Carl Christian:&nbsp;<em>Versuch eines Chiffernlexikons zu den G&ouml;ttinger, Vo&szlig;ischen, Schillerschen und Schlegel-Tieckschen Musenalmanachen.</em>&nbsp;Hamburg 1875.</td>\r\n</tr>\r\n<tr>\r\n<td>TB</td>\r\n<td>Thieme, Ulrich; Becker, Felix (Hg.):&nbsp;<em>Allgemeines Lexikon der bildenden K&uuml;nstler von der Antike bis zur Gegenwart</em>. 37 Bde. Leipzig 1907-1950. Neudruck Leipzig 1999.</td>\r\n</tr>\r\n</tbody>\r\n</table>")
	if err := dao.SaveRecord(model_lit_siglen); err != nil {
		fmt.Println(err)
	}

	model_lit_motto := models.NewRecord(collection)
	model_lit_motto.Set("Titel", "Literatur: Motto")
	model_lit_motto.Set("Text", "<p><strong>Zur Einstimmung</strong><br>Wie sie knallen die Peitschen! Hilf Himmel!<br>Journale! Calender!<br>Wagen an Wagen! Wieviel Staub und<br>wie wenig Gep&auml;ck!<br>(Xenien) Schillers Musenalmanach 1797 S. 260</p>")
	if err := dao.SaveRecord(model_lit_motto); err != nil {
		fmt.Println(err)
	}

	model_lit_kataloge := models.NewRecord(collection)
	model_lit_kataloge.Set("Titel", "Literatur: Kataloge")
	model_lit_kataloge.Set("Text", "<p><em>Die Almanache, Kalender und Taschenb&uuml;cher (1750-1860) der Landesbibliothek Coburg.</em> Wiesbaden 1970. BAUMG&Auml;RTEL, Ehrfried (Hrsg.).</p>\r\n<p><em>Almanache, Taschenb&uuml;cher, Taschenkalender.</em> Weimar 1967. Katalog der Sammlung der Th&uuml;ringischen Landesbibliothek Weimar mit 816 Nummern, 8 Abbildungen. MARWINSKI, Felicitas (Bearb.).</p>\r\n<p><em>Bild und Buch. Rheinbl&uuml;then, Moosrosen und Vergi&szlig;meinnicht: Taschenb&uuml;cher f&uuml;r Frauenzimmer von Bildung. Eine Studioausstellung in Zusammenarbeit mit der Badischen Landesbiliothek.</em> Staatliche Kunsthalle Karlsruhe 1995.</p>\r\n<p><em>F&uuml;rs sch&ouml;ne Geschlecht. Frauenalmanache zwischen 1800 und 1850.</em> Ausstellung der Universit&auml;t Bamberg in Zusammenarbeit mit der Staatsbibliothek Bamberg 12. November 1992 &ndash; 27. Februar 1993. Katalog und Ausstellung: Lydia Schieth. Bamberg o. J.</p>\r\n<p><em>Kalender? Ey, wie viel Kalender! Literarische Almanache zwischen Rokoko und Klassizismus.</em> Katalog und Ausstellung: York-Gothart MIX. Mit Beir&auml;gen von Karl-Heinz Hahn, Wolfgang Martens, John Mc Carthy, Regine Otto, Roger Paulin, Hartmut S&uuml;hrig und Herbert Zeman. Ausstellung im Zeughaus der Herzog August Bibliothek in Wolfenb&uuml;ttel vom 15. Juni bis 5. November 1986. Ausstellungskataloge der Herzog August Bibliothek Nr. 50. Wolfenb&uuml;ttel 1986.</p>\r\n<p><em>Kalender im Wandel der Zeiten. Eine Ausstellung der Badischen Landesbibliothek zur Erinnerung an die Kalenderreform durch Papst Gregor XIII. im Jahr 1582.</em> Austellungskatalog, herausgegeben von der Badischen Landesbibliothek Karlsruhe, unter Mitarbeit von Adrian Braunbehrens, R&uuml;diger Hannemann, Felix Heinzer, Jan Knopf, Gerhard R&ouml;mer, Ludwig Rohner, Blanka Tomanek. Karlsruhe, Bad. Landesbibliothek 1982.</p>\r\n<p><em>O sehet her! die allerliebsten Dingerchen. Friedrich R&uuml;ckert und der Almanach.</em> Eine Ausstellung der Bibliothek Otto Sch&auml;fer, des Stadtarchivs Schweinfurt, der St&auml;dtischen Sammlungen Schweinfurt und der R&uuml;ckert- Gesellschaft e.V. W&uuml;rzburg 2000.</p>\r\n<p><em>Taschenb&uuml;cher im 19. Jahrhundert.</em> Ausstellung im Schiller-Nationalmuseum Marbach zwischen November 1992 und Februar 1993. FALLBACHER, Karl-Heinz (Bearb.), Dt. Schillerges., Marbach am Neckar 1992.</p>\r\n<p><em>Wiener Kalender, Almanache und Taschenb&uuml;cher aus f&uuml;nf Jahrhunderten (1495-1977)</em> Wechselausstellung der Wiener Stadt- u. Landesbibliothek, Rathaus, Dezember 1976 &ndash; J&auml;nner 1977. Wiener Stadt- u. Landesbibliothek, f&uuml;r d. Inh. verantw.: Franz PATZER. Wien 1976.</p>")
	if err := dao.SaveRecord(model_lit_kataloge); err != nil {
		fmt.Println(err)
	}

	model_lit_forschung := models.NewRecord(collection)
	model_lit_forschung.Set("Titel", "Literatur: Forschungsliteratur")
	model_lit_forschung.Set("Text", "<p><strong>ANANIEVA, Anna; B&Ouml;CK, Dorothea; POMPE, Hedwig [Hrsg.]:</strong> <em>Geselliges Vergn&uuml;gen. Kulturelle Praktiken von Unterhaltung im langen 19. Jahrhundert.</em> Bielefeld 2011.</p>\r\n<p><strong>ANDERLE, Martin:</strong> <em>Wiener Lyrik im 18. Jahrhundert. Die Gedichte des &raquo;Wiener Musenalmanachs 1777-1796&laquo;.</em> Stuttgart 1996.</p>\r\n<p><strong>BENJAMIN, Walter:</strong> <em>Was die Deutschen lasen, w&auml;hrend ihre Klassiker schrieben.</em> [H&ouml;rst&uuml;ck] Ges. Werke hrsg. von Rolf Tiedeman u. Hermann Schweppenh&auml;user. Bd IV/1; Frankfurt/Main 1972, S. 641.</p>\r\n<p><strong>BOEHN, Max von:</strong> <em>Der Almanach.</em> in: <em>Das Antiquariat</em> 7 (1951), S. 3.</p>\r\n<p><strong>BUNZEL, Wolfgang:</strong> <em>Almanache und Taschenb&uuml;cher.</em> in: <strong>FISCHER, Ernst; HAEFS, Wilhelm; MIX, York-Gothart [Hrsg]:</strong> <em>Von Almanach bis Zeitung. Ein Handbuch der Medien in Deutschland 1700-1800.</em> M&uuml;nchen 1999, S. 24-35.</p>\r\n<p><strong>BUNZEL, Wolfgang:</strong> <em>Poetik und Publikation. Goethes Ver&ouml;ffentlichungen in Musenalmanachen und literarischen Taschenb&uuml;chern. Mit einer Bibliographie der Erst- und autorisierten Folgedrucke literarischer Texte Goethes im Almanach (1773-1832).</em> Weimar 1997.</p>\r\n<p><strong>CASSER, Paul:</strong> <em>Die westf&auml;lischen Musenalmanache und poetischen Taschenb&uuml;cher: ein Beitrag zur Geschichte Westfalens in der ersten H&auml;lfte des 19. Jahrhunderts.</em> Diss., M&uuml;nster 1928. Mikrofiche-Ausg.: Egelsbach 1992.</p>\r\n<p><strong>DICKENBERGER, Udo [Hrsg]:</strong> <em>Der Tod und die Dichter. Scherzgedichte in den Musenalmanachen um 1800. Eine Sammlung von 220 Spottgrabinschriften.</em> Hildesheim 1991.</p>\r\n<p><strong>D&Ouml;RING, Moritz</strong>: <em>Grenzen &uuml;berschreiten. Rezipienten-, Text-, Format- und Variantenwanderungen im</em> <em>&raquo;Taschenbuch zum geselligen Vergn&uuml;gen auf das Jahr 1823.&laquo;</em> Pfennig-Magazin zur Journalliteratur. Heft 3. Hannover 2018.</p>\r\n<p><strong>ENGELSING, Rolf:</strong> <em>Die Perioden der Lesergeschichte in der Neuzeit. Das statistische Ausma&szlig; und die soziokulturelle Bedeutung der Lekt&uuml;re.</em> in: AGB 10 (1970), S. 946-1002.</p>\r\n<p><strong>FISCHER, Bernhard:</strong> <em>Cottas &raquo;Morgenblatt f&uuml;r gebildete St&auml;nde&laquo; in der Zeit von 1807 bis 1823 und die Mitarbeit Therese Hubers.</em> in: <em>AGB</em> 43 (1995), S. 203-239.</p>\r\n<p><strong>FISCHER, Ernst; HAEFS, Wilhelm; MIX, York-Gothart:</strong> <em>Einleitung: Aufkl&auml;rung, &Ouml;ffentlichkeit und Medienkultur in Deutschland im 18. Jahrhundert.</em> in: <strong>FISCHER, Ernst; HAEFS, Wilhelm; MIX, York-Gothart [Hrsg]:</strong> <em>Von Almanach bis Zeitung. Ein Handbuch der Medien in Deutschland 1700-1800.</em> M&uuml;nchen 1999, S. 9-23.</p>\r\n<p><strong>FRIEDLAENDER, Max:</strong> <em>Gedichte von Goethe in Compositionen seiner Zeitgenossen.</em> Weimar 1896. Reprint Hildesheim 1975.</p>\r\n<p><strong>FRITSCH, Thomas Freiherr von:</strong> <em>Die Gothaischen Taschenb&uuml;cher, Hofkalender und Almanach.</em> Limburg an der Lahn 1968.</p>\r\n<p><strong>F&Uuml;RST, Rainer:</strong> <em>&raquo;F&uuml;r edle Weiber und M&auml;dchen.&laquo; Wilhelmine M&uuml;ller geb. Maisch, Verfasserin und F&ouml;rderin der Almanachliteratur um 1800. Ein Beitrag zur Verlagsgeschichte.</em> Heidelberg 1995.</p>\r\n<p><strong>GLADT, Karl:</strong> <em>Almanache und Taschenb&uuml;cher aus Wien.</em> Wien 1971</p>\r\n<p><strong>GLEISSNER, Stephanie; HUSIC, Mirela; KAMINSKI, Nicola; MERGENTHALER, Volker:</strong> <em>Optische Auftritte. Marktszenen in der medialen Konkurrenz von Journal-, Almanachs- und B&uuml;cherliteratur.</em> Hannover 2019 [= <em>Journalliteratur,</em> 2].</p>\r\n<p><strong>GOLDSCHMIDT, Arthur:</strong> <em>Goethe im Almanach.</em> Leipzig 1932.</p>\r\n<p><strong>GRANTZOW, Hans:</strong> <em>Geschichte des G&ouml;ttinger und des Vossischen Musenalmanachs [Kap. 1-4].</em> Diss., Berlin 1909.</p>\r\n<p><strong>GREILICH, Susanne:</strong> <em>Franz&ouml;sichsprachige Volksalmanache des 18. und 19. Jahrhunderts. Strukturen, Wandlungen, intertextuelle Bez&uuml;ge.</em> Heidelberg 2004.</p>\r\n<p><strong>HAEFS, Wilhelm:</strong> <em>Ein Kalender f&uuml;r die &raquo;mitleidigen Schwestern der Venus&laquo;? Die Literarisierung der Prostitution im Wiener &raquo;Taschenbuch f&uuml;r Grabennymphen auf das Jahr 1787&laquo;.</em> in: Jahrbuch der R&uuml;ckert- Gesellschaft e. V. 15 (2003), S. 101-110.</p>\r\n<p><strong>HAEFS, Wilhelm; MIX, York-Gothart:</strong> <em>Der Musenhort in der Provinz. Literarische Almanache in den Kronl&auml;ndern der &ouml;sterreichischen Monarchie im ausgehenden 18. und beginnenden 19. Jahrhundert.</em> in: <em>AGB</em> 27 (1986), S. 171‑194.</p>\r\n<p><strong>HAFERKORN, Hans J&uuml;rgen:</strong> <em>Der freie Schriftsteller. Eine literatur-soziologische Studie &uuml;ber seine Entstehung und Lage in Deutschland zwischen 1750 und 1800.</em> in: <em>AGB</em> 5 (1964), S. 523-713.</p>\r\n<p><strong>HAY, Gerhard:</strong> <em>Die Beitr&auml;ger des Voss&rsquo;schen Musenalmanaches. Ein Verzeichnis.</em> Hildesheim 1975.</p>\r\n<p><strong>HERZOG, Rudolph:</strong> <em>Die schlesischen Musenalmanache von 1773-1823.</em> Breslau 1912.</p>\r\n<p><strong>HISTORISCHES MUSEUM DER STADT WIEN:</strong> <em>Hieronymus L&ouml;schenkohl. 1753‑1807.</em> Wien 1959.</p>\r\n<p><strong>KAMINSKI, Nicola:</strong> <em>Die journalliterarische Leseszene im Spiegel des Modebilds. Modellversuch zur Wiener Zeitschrift 1816&ndash;1849</em>. Hannover 2022.</p>\r\n<p>&nbsp;</p>\r\n<p><strong>KLUSSMANN, Paul Gerhard; MIX, York-Gothart [Hrsg]:</strong> <em>Literarische Leitmedien. Almanach und Taschenbuch im kulturwissenschaftlichen Kontext.</em> Wiesbaden 1998 (Sammelband).</p>\r\n<p><strong>KOSSMANN, E. F.:</strong> <em>Der deutsche Musenalmanach 1833-1839.</em> Haag 1909.</p>\r\n<p><strong>LANCKORONSKA, Maria; R&Uuml;MANN, Arthur:</strong> <em>Geschichte der deutschen Taschenb&uuml;cher und Almanache aus der klassisch-romantischen Zeit.</em> M&uuml;nchen 1957. Neudruck Osnabr&uuml;ck 1985.</p>\r\n<p><strong>LIERES, Vita von:</strong> <em>Kalender und Almanache.</em> in: <em>Zeitschrift f&uuml;r B&uuml;cherfreunde</em> 18 (1926), S. 101-114.</p>\r\n<p><strong>LUDIN, Alfred:</strong> <em>Der schweizerische Musenalmanach &raquo;Alpenrosen&laquo; und seine Vorg&auml;nger (1780-1830).</em> Diss. Z&uuml;rich 1902.</p>\r\n<p><strong>L&Uuml;SEBRINK, Hans-J&uuml;rgen; MIX, York-Gothart u. a. (Hg.):</strong> <em>Franz&ouml;sische Almanachkultur im deutschen Sprachraum (1700-1815). Gattungsstrukturen, komparatistische Aspekte, Diskursformen.</em> G&ouml;ttingen 2013 [= <em>Deutschland und Frankreich im wissenschaftlichen Dialog,</em> 3].</p>\r\n<p><strong>MIX, York-Gothart:</strong> <em>Alamanach- und Taschenbuchkultur des 18. und 19. Jahrhunderts.</em> Wiesbaden 1996. [= Wolfenb&uuml;tteler Forschungen, Bd. 69]</p>\r\n<p><strong>MIX, York-Gothart:</strong> <em>Die deutschen Musenalmanache des 18. Jahrhunderts.</em> M&uuml;nchen 1987.</p>\r\n<p><strong>MIX, York-Gothart:</strong> <em>Geselligkeitskultur, Gattungskonvention und Publikumsinteresse. Zur Intention und Funktion von C. M. Wielands und J. W. v. Goethes &raquo;Taschenbuch auf das Jahr 1804&laquo; und O. J. Bierbaums &raquo;Modernem Musen-Almanach&laquo;.</em> in: <em>Jahrbuch des Wiener Goethe-Vereins</em> 97/98 (1993), S. 35-45.</p>\r\n<p><strong>OBENAUS, Sibylle:</strong> <em>Die deutschen allgemeinen kritischen Zeitschriften in der ersten H&auml;lfte des 19. Jahrhunderts.</em> in: <em>AGB</em> 14 (1974), S. 2-122.</p>\r\n<p><strong>PEPERKORN, G&uuml;nter:</strong> <em>Dieses ephemerische Werckchen: Georg Christoph Lichtenberg und der G&ouml;ttinger Taschen Calender.</em> G&ouml;ttingen [St&auml;dt. Museum] 1992.</p>\r\n<p><strong>PISSIN, Raimund:</strong> <em>Almanache der Romantik.</em> Berlin-Zehlendorf 1910.</p>\r\n<p><strong>PFEIFFER, Emil:</strong> <em>Bibliographie der Schillerschen Musenalmanache 1796-1800.</em> in: <em>Jahresbericht des Schw&auml;bischer Schillerverein.</em> Marbach 1916, S. 35-48</p>\r\n<p><strong>PFISTER, Karl:</strong> <em>Das Prinzip der Gedichtanordnung in Schillers Musenalmanachen 1796/1800.</em> Diss., Bern 1937.</p>\r\n<p><strong>PR&Uuml;SENER, Marlies:</strong> <em>Lesegesellschaften im 18. Jahrhundert.</em> in: <em>AGB</em> 13 (1973), S. 371-594.</p>\r\n<p><strong>PRUTZ, Robert:</strong> <em>Der G&ouml;ttinger Dichterbund. Zur Geschichte der deutschen Literatur.</em> Leipzig 1841.</p>\r\n<p><strong>PRUTZ, Robert:</strong> <em>Neue Schriften. Zur deutschen Literatur- und Kulturgeschichte.</em> 2 Bde.; Halle 1847. bes.: Bd. I. S. 105-165: <em>Die Musenalmanache und Taschenb&uuml;cher in Deutschland.</em></p>\r\n<p><strong>RAABE, Paul:</strong> <em>Zeitschriften und Almanache.</em> in: <em>Buchkunst und Literatur in Deutschland 1750 bis 1850.</em> Herausgegeben von Ernst Hauswedell und Christian Voigt. Hamburg 1977, Bd. 1. S. 145-195 [mit ausf&uuml;hrlichem Abbildungsteil in Bd. 2. S. 108-140].</p>\r\n<p><strong>REDLICH, Carl Christian:</strong> <em>Versuch eines Chiffernlexikons zu den G&ouml;ttinger, Vo&szlig;ischen, Schillerschen und Schlegel-Tieckschen Musenalmanachen.</em> Hamburg 1875.</p>\r\n<p><strong>ROMMEL, Otto:</strong> <em>Der Wiener Musenalmanach.</em> in: <em>Euphorion</em> 6. Erg&auml;nzungsheft, 1906.</p>\r\n<p><strong>SCHR&Ouml;DER, Rolf:</strong> <em>Zur Struktur des &raquo;Taschenbuchs&laquo; im Biedermeier.</em> in: <em>Germanisch-Romanische Monatsschrift</em> 41 (1960), S. 442-448.</p>\r\n<p><strong>SCHWERDTFEGER, Walter:</strong> <em>Die litteraturhistorische Bedeutung der Schillerschen Musenalmanache 1796-1800.</em> Leipzig 1899.</p>\r\n<p><strong>SEYFFERT, Wolfgang:</strong> <em>Schillers Musenalmanache.</em> Berlin 1913.</p>\r\n<p><strong>SKREB, Zdenko:</strong> <em>Das Epigramm in deutschen Musenalmanachen und Taschenb&uuml;chern um 1800.</em> Wien, 1977 [= <em>&Ouml;sterreichische Akademie der Wissenschaften, Philosophisch-historische Klasse, Sitzungsberichte,</em> 331].</p>\r\n<p><strong>SKREB, Zdenko:</strong> <em>Gattungsdominanz im deutschsprachigen literarischen Taschenbuch oder vom Sieg der Erz&auml;hlprosa,</em> Wien 1986 [= <em>&Ouml;sterreichische Akademie der Wissenschaften, Philosophisch-historische Klasse, Sitzungsberichte,</em> 471].</p>\r\n<p><strong>STEIG, Reinhold:</strong> <em>Ueber den G&ouml;ttingischen Musen-Almanach f&uuml;r das Jahr 1803.</em> in: <em>Euphorion</em> 2 (1895), S. 312-323</p>\r\n<p><strong>STOLPE, Heinz:</strong> <em>Zeitschriften und Almanache der deutschen Klassik.</em> Weimar 1959.</p>\r\n<p><strong>WILLNAT, Elisabeth:</strong> <em>Johann Christian Dieterich. Ein Verlagsbuchh&auml;ndler und Drucker in der Zeit der Aufkl&auml;rung.</em> in: <em>AGB</em> 39 (1993), S. 1-254.</p>\r\n<p><strong>WITTMANN, Reinhard:</strong> <em>Der Verleger Johann Friedrich Weygand in Briefen des G&ouml;ttinger Hains.</em> in: AGB 10 (1970), S. 319-343.</p>\r\n<p><strong>ZUBER, Margarete:</strong> <em>Die deutschen Musenalmanache und sch&ouml;ngeistigen Taschenb&uuml;cher des Biedermeier 1815- 1848.</em> in: <em>AGB</em> 1 (1958), S. 398-489.</p>")
	if err := dao.SaveRecord(model_lit_forschung); err != nil {
		fmt.Println(err)
	}

	model_lit_graphik := models.NewRecord(collection)
	model_lit_graphik.Set("Titel", "Literatur: Buchkunst")
	model_lit_graphik.Set("Text", "<p><strong>BARGE, Hermann:</strong> <em>Geschichte der Buchdruckerkunst.</em> Leipzig 1940.</p>\r\n<p><strong>BAUER, Jens-Heiner:</strong> <em>Daniel Nikolaus Chodowiecki. Das druckgraphische Werk. Die Sammlung Wilhelm Burggraf zu Dohna-Schlobitten. Ein Bildband mit 2340 Abbildungen in Erg&auml;nzung zum Werkverzeichnis von Wilhelm Engelmann.</em> Hannover 1982.</p>\r\n<p><strong>DORN, Wilhelm:</strong> <em>Meil-Bibliographie. Verzeichnis der von dem Radierer Johann Wilhelm Meil illustrierten B&uuml;cher und Almanache</em> Berlin 1928.</p>\r\n<p><strong>FOCKE Rudolf [Hg.]:</strong> <em>Chodowiecki und Lichtenberg. Daniel Chodowiecki&rsquo;s Monatskupfer zum &raquo;G&ouml;ttinger Taschen Calender&laquo; nach Georg Christoph Lichtenberg&rsquo;s Erkl&auml;rungen (1778-1783), mit einer kunst- und litterargeschichtlichen Einleitung.</em> Leipzig 1901.</p>\r\n<p><strong>FORSTER-HAHN Franziska:</strong> <em>Johann Heinrich Ramberg als Karikaturist und Satiriker.</em> Diss. Univ. Bonn 1963. [o. O.] [o. J.] [= <em>Sonderdruck aus Hann. Geschichtsbl&auml;ttern,</em> NF 17 (1963)].</p>\r\n<p><strong>HALDENWANG Hasso von:</strong> <em>Christian Haldenwang, Kupferstecher (1770-1831).</em> Diss. Johann-Wolfgang-Goethe-Univ. Frankfurt am Main 1995, Frankfurt am Main 1997 [= <em>Frankfurter Fundamente der Kunstgeschichte,</em> 14].</p>\r\n<p><strong>HAUSWEDELL, Ernst L.; VOIGT, Christian [Hrsg.]:</strong> <em>Buchkunst und Literatur in Deutschland 1750 bis 1850.</em> 2 Bde., Hamburg 1977.</p>\r\n<p><strong>KO&Scaron;ENINA, Alexander [Hrsg.]:</strong> <em>Literatur &mdash; Bilder. Johann Heinrich Ramberg als Buchillustrator der Goethezeit.</em> Hannover 2013.</p>\r\n<p><strong>LANCKORONSKA, Maria; Oehler, Richard:</strong> <em>Die Buchillustration dex XVIII. Jahrhunderts in Deutschland, &Ouml;sterreich und der Schweiz.</em> 3 Bde., Leipzig 1932-1934.</p>\r\n<p><strong>RODENBERG, J.:</strong> <em>Geschichte der Illustration von 1800 bis heute.</em> in: <strong>LEIH, G. [Hg.]:</strong> <em>Handbuch der Bibliothekswissenschaft.</em> 2. Aufl. Stuttgart 1950, Bd. 1.</p>\r\n<p><strong>RHEIN, Adolf:</strong> <em>Die fr&uuml;hen Verlagseinb&auml;nde. Eine technische Entwicklung 1735-1850.</em> in: <em>Gutenberg-Jahrbuch,</em> Mainz 1962, S. 519-532.</p>\r\n<p><strong>R&Uuml;MANN, Arthur:</strong> <em>Das illustrierte Buch des XIX. Jahrhunderts in England, Frankreich und Deutschland 1790-1860.</em> Nachdruck der Ausgabe des Insel Verlages 1930, Osnabr&uuml;ck 1975.</p>\r\n<p><strong>R&Uuml;MANN, Arthur:</strong> <em>Die illustrierten deutschen B&uuml;cher des 18. Jahrhunderts.</em> Stuttgart 1927.</p>\r\n<p><strong>R&Uuml;MANN, Arthur:</strong> <em>Die illustrierten deutschen B&uuml;cher des 19. Jahrhunderts.</em> Stuttgart 1926.</p>\r\n<p><strong>R&Uuml;MANN, Arthur:</strong> <em>Das deutsche illustrierte Buch des XVIII. Jahrhunderts.</em> Stra&szlig;burg 1931 [= <em>Studien zur deutschen Kunstgeschichte,</em> Heft 282].</p>\r\n<p><em>Sammlung Hogarthscher Kupfer-Stiche.</em> Neue wohlfeile Ausg., G&ouml;ttingen [o. J.].</p>\r\n<p><strong>SCHUMACHER, Doris:</strong> <em>Kupfer und Poesie. Die Illustrationskunst um 1800 im Spiegel der zeitgen&ouml;ssischen deutschen Kritik.</em> K&ouml;ln 2000 [= <em>Pictura et Poesis,</em> 13].</p>\r\n<p><strong>SHESGREEN, Sean (Hg.):</strong> <em>Engravings by Hogarth. 101 Prints.</em> New York 1973.</p>\r\n<p><strong>STUBBE, Wolf:</strong> <em>Illustrationen und Illustratoren.</em> in: <strong>HAUSWEDELL, Ernst; VOIGT, Christian [Hg.]:</strong> <em>Buchkunst und Literatur in Deutschland 1750 bis 1850.</em> Bd. 1., Hamburg 1977, S. 58-144 [mit ausf&uuml;hrlichem Abbildungsteil Bd. 2. S. 49-106]</p>\r\n<p><strong>STUTTMANN, Ferdinand:</strong> <em>Johann Heinrich Ramberg.</em> M&uuml;nchen 1929.</p>")
	if err := dao.SaveRecord(model_lit_graphik); err != nil {
		fmt.Println(err)
	}

	model_lit_nachschlagewerke := models.NewRecord(collection)
	model_lit_nachschlagewerke.Set("Titel", "Literatur: Nachschlagewerke")
	model_lit_nachschlagewerke.Set("Text", "<p><em>Allgemeine Deutsche Biographie.</em> Hg. v. der Historischen Commission bei der K&ouml;niglichen Akademie der Wissenschaften. 55 Bde., Leipzig 1875-1910. Sigle: ADB.</p>\r\n<p><em>Brockhaus&rsquo; Conversations-Lexikon. Allgemeine deutsche Real-Encyklop&auml;die.</em> 13. vollst. umgearb. Aufl., Leipzig 1883-1888. Sigle: Brockh 13.</p>\r\n<p><strong>EYMER, Wilfried:</strong> <em>Eymers Pseudonymen Lexikon. Realnamen und Pseudonyme in der deutschen Literatur.</em> Bonn 1997. Sigle: E.</p>\r\n<p><strong>FISCHER, Bernhard:</strong> <em>Der Verleger Johann Friedrich Cotta. Chronologische Verlagsbibliographie 1787-1814. Aus den Quellen bearbeitet.</em> 3 Bde., M&uuml;nchen 2003.</p>\r\n<p><strong>GOEDEKE, Karl u. a.:</strong> <em>Grundri&szlig; zur Geschichte der deutschen Dichtung.</em> 13 Bde., 2. Aufl, Dresden 1884 ff.; (Bd. IV/1-4 in der dritten neubearbeiteten Aufl., Dresden 1896-1913. Neudruck Berlin 1955; Bd. IV/5 in der ersten Aufl. D&uuml;sseldorf 1957-1960; au&szlig;erdem die &bdquo;Neue Folge&ldquo;, N. F. Bd. I, Berlin 1955ff.). Sigle: Goed</p>\r\n<p><strong>GOLDSCHMIDT, Arthur:</strong> <em>Goethe im Almanach.</em> Leipzig 1932. Sigle: G.</p>\r\n<p><strong>GORZNY, Willy:</strong> <em>Deutscher Biographischer Index.</em> Bearb. v. Hans-Albrecht Koch, Uta Koch und Angelika Koller, 4 Bde., M&uuml;nchen 1986 [Register zu: <strong>GORZNY, Willy [Hg.]</strong> <em>Deutsches Biographisches Archiv.</em> M&uuml;nchen 1985]. Sigle: DBI.</p>\r\n<p><strong>GRAND-CARTERET, John:</strong> <em>Les Almanachs Francais. Bibliographie &ndash; Iconographie 1600-1895.</em> Paris 1896. Sigle: G-C.</p>\r\n<p><strong>HAYN, Hugo; GOTENDORF, Alfred N. (Hg):</strong> <em>Bibliotheca Germanorum Erotica &amp; Curiosa. Verzeichnis der gesamten deutschen erotischen Literatur mit Einschlu&szlig; der &Uuml;bersetzungen, nebst Beif&uuml;gung der Originale.</em> 9 Bde., Unver&auml;nd. Nachdr. d. 3. ungemein verm. Aufl. Hanau [o. J.], Hanau 1968. Sigle: H.-G.</p>\r\n<p><strong>HIRSCHBERG, Leopold:</strong> <em>Der Taschengoedeke. Bibliographie deutscher Erstausgaben.</em> M&uuml;nchen 1970. Sigle: Goed (H).</p>\r\n<p><strong>HOEFER, Jean Chr&eacute;tien Ferdinand (Hg.):</strong> <em>Nouvelle biographie g&eacute;n&eacute;rale.</em> 46 Bde. Paris 1851-1878.</p>\r\n<p><strong>HOLZMANN, Michael; BOHATTA, Hanns:</strong> <em>Deutsches Pseudonymen-Lexikon. Aus den Quellen bearbeitet von Michael Holzmann und Hanns Bohatta.</em> Hildesheim 1970. Sigle: H/B.</p>\r\n<p><strong>K&Ouml;HRING, Hans (Bearb./Hg.):</strong> <em>Bibliographie der Almanache, Kalender und Taschenb&uuml;cher f&uuml;r die Zeit von ca. 1750-1860.</em> Hamburg 1929. Neudruck Bad Karlshafen 1987. Sigle: K.</p>\r\n<p><strong>LANCKORONSKA, Maria; R&Uuml;MANN, Arthur:</strong> <em>Geschichte der deutschen Taschenb&uuml;cher und Almanache aus der klassisch-romantischen Zeit.</em> M&uuml;nchen 1957. Neudruck Osnabr&uuml;ck 1985. Sigle: L/R.</p>\r\n<p><strong>R&Uuml;MANN, Arthur:</strong> <em>Die illustrierten deutschen B&uuml;cher des 18. Jahrhunderts.</em> Stuttgart 1927. Sigle: R&uuml;.</p>\r\n<p><strong>REDLICH, Carl Christian:</strong> <em>Versuch eines Chiffernlexikons zu den G&ouml;ttinger, Vo&szlig;ischen, Schillerschen und Schlegel-Tieckschen Musenalmanachen.</em> Hamburg 1875. Sigle: Re.</p>\r\n<p><strong>THIEME, Ulrich; BECKER, Felix [Hg.]:</strong> <em>Allgemeines Lexikon der bildenden K&uuml;nstler von der Antike bis zur Gegenwart.</em> 37 Bde., Leipzig 1907-1950. Neudruck Leipzig 1999. Sigle: T/B.</p>\r\n<p><strong>ZIEGLER, Konrad; SONTHEIMER, Walther [Hg.]:</strong> <em>Der kleine Pauly. Lexikon der Antike.</em> 5 Bde., M&uuml;nchen 1979. Sigle: P.</p>")
	if err := dao.SaveRecord(model_lit_nachschlagewerke); err != nil {
		fmt.Println(err)
	}

	model_danksagungen := models.NewRecord(collection)
	model_danksagungen.Set("Titel", "Danksagungen")
	model_danksagungen.Set("Text", "<div><p>Der bibliographische Auftrieb auf die Musenalm bedarf der Unterst&uuml;tzung durch ihre Nutzer und insbesondere durch die Besitzer seltener und wenig bekannter, kaum auffindbarer Almanache und Taschenb&uuml;cher des gew&auml;hlten Zeitraumes von etwa 1750 bis 1870. Sie helfen uns durch Hinweise, Leihgaben auf kurze Frist, sowie durch Benennung von Fehlern und Unstimmigkeiten, die uns unterlaufen m&ouml;gen, sich aber auch aus Eigen- und Abarten untersuchter Exemplare herleiten k&ouml;nnen.</p><p>F&uuml;r viele geleistete Hilfen danken wir:</p><p>Frau <strong>Susanne Koppel</strong><br>Antiquariat Susanne Koppel<br>Parkallee 4<br>20144 Hamburg<br><a href=\"http://www.antiquariat-koppel.de\" target=\"_blank\" rel=\"noreferrer noopener\">www.antiquariat-koppel.de</a><br><a href=\"mailto:info@antiquariat-koppel.de\">info@antiquariat-koppel.de</a></p><p>Herrn <strong>Thomas Rezek</strong><br>Antiquariat Thomas Rezek<br>Amalienstra&szlig;e 63<br>80799 M&uuml;nchen<br><a href=\"http://www.a-rezek.de\" target=\"_blank\" rel=\"noreferrer noopener\">www.a-rezek.de</a><br><a href=\"mailto:arezek@web.de\">arezek@web.de</a></p><p>Herrn <strong>G&uuml;nther Trauzettel-Loibl</strong><br>Antiquariat Trauzettel<br>Haum&uuml;hle 8<br>52223 Stolberg<br><a href=\"http://www.antiquariat-trauzettel.de\" target=\"_blank\" rel=\"noreferrer noopener\">www.antiquariat-trauzettel.de</a><br><a href=\"mailto:antiquariat.trauzettel@t-online.de\">antiquariat.trauzettel@t-online.de</a></p><p>Herrn <strong>Uwe Turszynski</strong><br>Antiquariat Turszynski<br>Herzogstra&szlig;e 66<br>80803 M&uuml;nchen<br><a href=\"http://www.turszynski.de\" target=\"_blank\" rel=\"noreferrer noopener\">www.turszynski.de</a><br><a href=\"mailto:antiquariat@turszynski.de\">antiquariat@turszynski.de</a></p><p>Herrn <strong>Dieter Zipprich</strong><br>Antiquariat Zipprich<br>Karolinenstra&szlig;e 18<br>96049 Bamberg<br><a href=\"mailto:antiquariat.zipprich@freenet.de\">antiquariat.zipprich@freenet.de</a></p><p>Frau Mag. <strong>Rita Robosch</strong><br>Matthaeus Truppe Buchhandlung &amp; Antiquariat<br>Stubenberggasse 7<br>A-8010 Graz<br>Austria<br><a href=\"mailto:truppe@aon.at\">truppe@aon.at</a></p></div>")
	if err := dao.SaveRecord(model_danksagungen); err != nil {
		fmt.Println(err)
	}

	model_impressum := models.NewRecord(collection)
	model_impressum.Set("Titel", "Impressum")
	model_impressum.Set("Text", "<div><p><strong>Telemedienanbieter im Sinne des &sect; 5 TMG:</strong><br>Theodor Springmann Stiftung<br>Hirschgasse 2<br>69120 Heidelberg<br><br>Telefon +49 6221 436235<br>Email&nbsp;<a href=\"mailto:info@theodor-springmann-stiftung.de\">info@theodor-springmann-stiftung.de</a></p><p><strong>Rechtsform und Sitz:</strong><br>Die Theodor Springmann Stiftung ist eine rechtsf&auml;hige Stiftung b&uuml;rgerlichen Rechts. Sitz der Stiftung ist Heidelberg.</p><p><strong>Vorstand der Theodor Springmann Stiftung:</strong><br>Dr. Randolf Straky (Pr&auml;sident)</p><p><strong>Gesch&auml;ftsf&uuml;hrung:</strong><br>Dr. Randolf Straky<br>Theodor Springmann Stiftung<br>Hirschgasse 2<br>69120 Heidelberg</p><p><strong>Zust&auml;ndige Aufsichtsbeh&ouml;rde:</strong><br>Regierungspr&auml;sidium Karlsruhe<br>Schlossplatz 1-3<br>76131 Karlsruhe<br><br><a href=\"https://rp.baden-wuerttemberg.de/Themen/Stiftung/Seiten/Ansprechpartner.aspx\">Ansprechpartner</a></p><p><strong>Redaktionelle Verantwortung nach &sect; 55 Abs. 2 RStV:</strong><br>Adrian Braunbehrens<br>Theodor Springmann Stiftung<br>Hirschgasse 2<br>69120 Heidelberg</p></div>")
	if err := dao.SaveRecord(model_impressum); err != nil {
		fmt.Println(err)
	}

	model_datenschutz := models.NewRecord(collection)
	model_datenschutz.Set("Titel", "Datenschutz")
	model_datenschutz.Set("Text", "<div><h2>Pr&auml;ambel</h2><p>Diese Datenschutzerkl&auml;rung informiert Sie &uuml;ber die Art, den Umfang und den Zweck der personenbezogenen Daten, die im Rahmen dieser Onlinepr&auml;senz von Ihnen erhoben und von uns verarbeitet werden, sowie die Ihnen zustehenden Rechte.<br>Personenbezogene Daten sind alle Informationen, die sich auf eine identifizierte oder identifizierbare nat&uuml;rliche Person beziehen. Als identifizierbar wird eine nat&uuml;rliche Person angesehen, die direkt oder indirekt identifiziert werden kann. Im Hinblick auf weitere verwendete Begrifflichkeiten verweisen wir auf die Definitionen der Datenschutz-Grundverordnung (DSGVO), Artikel 4.<br>Erfolgt die Verarbeitung personenbezogener Daten auf Grundlage des Art. 6 Abs. 1 lit. f DSGVO, so besteht unser berechtigtes Interesse in der Erf&uuml;llung unseres satzungsgem&auml;&szlig;en Stiftungszwecks. Zweck der Stiftung ist es, in gemeinn&uuml;tziger Weise Wissenschaft und Kunst, V&ouml;lkerverst&auml;ndigung und Entwicklungshilfe zu f&ouml;rdern und zum menschlichen Selbstverst&auml;ndnis sowie zum Erkennen und Lindern strukturell bedingter Not und Bed&uuml;rftigkeit von Menschen in aller Welt beizutragen.</p>\r\n<h2>Verantwortlicher</h2>\r\n<p>Verantwortlich im Sinne von Art. 4 Abs. 7 DSGVO f&uuml;r die Verarbeitung personenbezogener Daten ist:<br>Theodor Springmann Stiftung<br>Hirschgasse 2<br>69120 Heidelberg<br><br><a href=\"mailto:info@theodor-springmann-stiftung.de\">info@theodor-springmann-stiftung.de</a></p>\r\n<h2>Automatisch bei Nutzung unserer Onlinepr&auml;senz erfasste Daten</h2>\r\n<p>Mit der Nutzung unserer Onlinepr&auml;senz werden automatisch personenbezogene und allgemeine Daten vom Nutzer &uuml;bermittelt und von uns erfasst und gespeichert. Die von uns erhobenen Daten umfassen:</p>\r\n<ol>\r\n<li>die IP-Adresse des ans Internet angeschlossenen Netzwerkger&auml;ts (Computer oder Router des Nutzers),</li>\r\n<li>den verwendeten Browsertyp und dessen Version,</li>\r\n<li>das verwendete Betriebssystem und dessen Version,</li>\r\n<li>die Internetseite, die den Nutzer auf unsere Onlinepr&auml;senz verweist (Referrer),</li>\r\n<li>die Unterwebseiten, die der Nutzer auf unserer Onlinepr&auml;senz aufruft,</li>\r\n<li>das Datum und die Uhrzeit des Aufrufs,</li>\r\n<li>den Internet-Service-Provider des Nutzers,</li>\r\n<li>sonstige &auml;hnliche Daten.</li>\r\n</ol>\r\n<p>Die Theodor Springmann Stiftung zieht aus diesen Daten keine R&uuml;ckschl&uuml;sse auf die betroffenen Nutzer. Diese Informationen werden ben&ouml;tigt, um</p>\r\n<ol>\r\n<li>die vom Nutzer angeforderten Inhalte korrekt auszuliefern,</li>\r\n<li>technische Probleme zu diagnostizieren, unsere IT-Systeme vor Angriffen zu sch&uuml;tzen und im Falle eines Angriffs den zust&auml;ndigen Beh&ouml;rden notwendige Informationen zur Strafverfolgung bereitstellen zu k&ouml;nnen,</li>\r\n<li>die Inhalte und die Bedienung unserer Onlinepr&auml;senz zu optimieren.</li>\r\n</ol>\r\n<p>Die oben genannten Verarbeitungszwecke werden von Subsystemen erf&uuml;llt, die unabh&auml;ngig von anderen Subsystemen ihre jeweils eigene Kopie der vom Nutzer erhobenen Daten erhalten, verarbeiten und l&ouml;schen. Dabei werden den Subsystemen lediglich die zur Ausf&uuml;hrung ihrer Aufgabe erforderlichen Daten &uuml;bergeben. Alle Subsysteme befinden sich auf Servern der Theodor Springmann Stiftung.</p>\r\n<h3>Auslieferung der vom Nutzer angeforderten Inhalte</h3>\r\n<p>Mit der Anforderung einer Seite dieser Onlinepr&auml;senz &uuml;bermittelt der Browser des Nutzers automatisch die IP-Adresse seines ans Internet angeschlossenen Netzwerkger&auml;ts (Computer oder Router) und andere Daten. Ohne diese IP-Adresse k&ouml;nnen die Inhalte dieser Online-Pr&auml;senz nicht an den Nutzer zur&uuml;ckgesendet werden. Daher ist die Speicherung und Verarbeitung der IP-Adresse f&uuml;r den Betrieb dieser Onlinepr&auml;senz notwendig. Ebenfalls &uuml;bermittelte Daten &uuml;ber den verwendeten Browser usw. werden gegebenenfalls verwendet, um die Anzeige der Webseite an das Anzeigeger&auml;t oder den Browser anzupassen. Die zum Zwecke der Auslieferung von Inhalten gesammelten Daten werden direkt nach &Uuml;bertragung der Inhalte vom Subsystem gel&ouml;scht. Die Verarbeitung dieser Daten erfolgt auf der Grundlage eines berechtigten Interesses nach Art. 6 Abs. 1 lit. f DSGVO.</p>\r\n<h3>Diagnose und Schutz der IT-Systeme</h3>\r\n<p>Daneben erfolgt eine Speicherung der oben genannten Daten in Diagnose-Protokollen. Die Protokollierung ist notwendig, um etwaige technische Probleme analysieren oder Angriffe erkennen und abwehren zu k&ouml;nnen. Bei Angriffen auf unsere IT-Systeme &uuml;bergeben wir die Diagnose-Protokolle den zust&auml;ndigen Strafverfolgungsbeh&ouml;rden. Die Diagnose-Protokolle werden nach 14 Tagen automatisch gel&ouml;scht. Die Speicherung und Verarbeitung dieser Daten erfolgt auf Grundlage eines berechtigten Interesses nach Art. 6 Abs. 1 lit. f DSGVO.</p>\r\n<h3>Analyse zur Optimierung der Inhalte und der Bedienung</h3>\r\n<p>Die oben genannten Daten werden von einem Webanalyseprogramm verarbeitet. Dabei werden die Daten pseudonymisiert und aggregiert, so dass sie einzelnen Nutzern nicht mehr zugeordnet werden k&ouml;nnen. Das Analyseprogramm wird auf einem von uns betriebenen Server ausgef&uuml;hrt, daher werden keine Daten an Dritte &uuml;bermittelt. Der Zweck dieser Erfassung besteht in der Analyse der Nutzung unseres Angebots, durch die eine Verbesserung unserer Webseiten, der Inhalte und der Bedienung erm&ouml;glicht wird. Das Analyseprogramm speichert die anonymisierten Daten ohne zeitliche Begrenzung. Die Verarbeitung dieser Daten erfolgt auf der Grundlage eines berechtigten Interesses (Art. 6 Abs. 1 lit. f DSGVO).</p>\r\n<h2>Sonstige Daten</h2>\r\n<p>S&auml;mtliche Daten, die mit dem Aufruf einer Webseite von unserer Onlinepr&auml;senz an den Nutzer &uuml;bertragen werden, werden von unseren Servern bereitgestellt. Wir nutzen kein von Dritten bereitgestelltes Content Delivery Network (CDN), um Teile unserer Onlinepr&auml;senz (z. B. Javascript- oder Webfont-Dateien) von dort an den Nutzer zu &uuml;bermitteln. Wir verlinken auch nicht in soziale Netzwerke. Insofern k&ouml;nnen von Dritten keine personenbezogenen Daten &uuml;ber unsere Onlinepr&auml;senz erhoben werden, es sei denn, die Datenerfassung durch Dritte basiert auf einer gesetzlichen Grundlage und/oder wurde beh&ouml;rdlich angeordnet.<br>In unseren Inhalten k&ouml;nnen sich jedoch Hyperlinks (&bdquo;Links&ldquo;) auf fremde Onlineangebote befinden. Mit dem Anklicken eines solchen Links verl&auml;sst der Nutzer unsere Onlinepr&auml;senz und damit den Geltungsbereich dieser Datenschutzerkl&auml;rung.</p>\r\n<h2>Erhobene Daten bei Kontakt</h2>\r\n<p>Unsere Onlinepr&auml;senz h&auml;lt unter anderem aufgrund gesetzlicher Vorgaben verschiedene M&ouml;glichkeiten bereit, mit uns in Kontakt zu treten. Sofern eine betroffene Person Kontakt zu uns aufnimmt, werden die von der betroffenen Person &uuml;bermittelten personenbezogenen Daten automatisch gespeichert (z. B. Telefonnummer, E-Mail-Adresse). Diese Angaben werden zum Zwecke der Bearbeitung sowie f&uuml;r sich m&ouml;glicherweise anschlie&szlig;ende Kommunikation von uns gespeichert. Sobald die Speicherung dieser Daten nicht mehr erforderlich ist und keine gesetzlichen Archivierungsgr&uuml;nde vorliegen, werden die Daten gel&ouml;scht. Die Erforderlichkeit wird alle zwei Jahre &uuml;berpr&uuml;ft. Es erfolgt keine Weitergabe dieser personenbezogenen Daten an Dritte.<br>Von betroffenen Personen &uuml;bermittelte Informationen zur Bearbeitung einer Kontaktanfrage werden gem&auml;&szlig; Art. 6 Abs. 1 lit. b (vertragliche und vorvertragliche Beziehungen) oder lit. f (andere Anfragen) der DSGVO verarbeitet.</p>\r\n<h2>Speicherung und L&ouml;schung von personenbezogenen Daten</h2>\r\n<p>Personenbezogene Daten werden von uns f&uuml;r die Dauer der entsprechenden gesetzlichen Aufbewahrungspflichten oder f&uuml;r den zur Vertragserf&uuml;llung erforderlichen Zeitraum gespeichert. Liegen solche Gr&uuml;nde nicht vor, werden personenbezogene Daten f&uuml;r den Zeitraum, der zur Erreichung des Speicherzwecks erforderlich ist, verarbeitet und gespeichert, sofern in dieser Datenschutzerkl&auml;rung nicht ausdr&uuml;cklich anderes angegeben ist. Sobald die Daten f&uuml;r ihre Zweckbestimmung nicht mehr erforderlich sind oder der Speicherungszweck entf&auml;llt, werden sie nach Ma&szlig;gabe der gesetzlichen Vorschriften gel&ouml;scht oder in ihrer Verarbeitung eingeschr&auml;nkt.<br>Sollten die Daten nicht gel&ouml;scht werden d&uuml;rfen, weil sie einer gesetzlichen Aufbewahrungspflicht unterliegen, wird deren Verarbeitung durch Archivierung eingeschr&auml;nkt. Die Daten werden erst nach Ablauf der gesetzlich vorgeschriebenen Speicherfrist gel&ouml;scht.<br>Alle sechs Monate wird routinem&auml;&szlig;ig gepr&uuml;ft, ob der Speicherungszweck weggefallen bzw. die Aufbewahrungspflicht abgelaufen ist. Anschlie&szlig;end wird gegebenenfalls die L&ouml;schung durchgef&uuml;hrt.</p>\r\n<h2>Rechte der betroffenen Personen</h2>\r\n<p>Unter den angegebenen Kontaktdaten k&ouml;nnen betroffene Personen jederzeit nachgenannte Rechte aus&uuml;ben. Eine betroffene Person kann von uns verlangen, dass</p>\r\n<ul>\r\n<li>Auskunft &uuml;ber sie betreffende, bei uns gespeicherte Daten und deren Verarbeitung erteilt wird (Art. 15 DSGVO),</li>\r\n<li>sie betreffende unrichtige personenbezogene Daten berichtigt werden (Art. 16 DSGVO),</li>\r\n<li>sie betreffende, bei uns gespeicherte personenbezogene Daten gel&ouml;scht werden (Art. 17 DSGVO),</li>\r\n<li>die Verarbeitung sie betreffender, bei uns gespeicherter Daten, die aufgrund gesetzlicher Vorschriften oder anderer Gr&uuml;nde nicht gel&ouml;scht werden d&uuml;rfen, eingeschr&auml;nkt wird (Art. 18 DSGVO),</li>\r\n<li>sie betreffende Daten &uuml;bertragbar sind, soweit sie in die Datenverarbeitung eingewilligt oder einen Vertrag mit uns geschlossen hat (Art. 20 DSGVO),</li>\r\n<li>sie betreffende Daten nach einem Widerspruch nicht weiter von uns verarbeitet werden (Art. 21 DSGVO).</li>\r\n</ul>\r\n<p>Eine betroffene Person hat ferner jederzeit das Recht,</p>\r\n<ul>\r\n<li>eine erteilte Einwilligung in die Erhebung und Verarbeitung personenbezogener Daten nach Art. 6 Abs. 1 lit. a oder Art. 9 Abs. 2 lit. a DSGVO f&uuml;r die Zukunft zu widerrufen (Art. 7 Abs. 3 DSGVO),</li>\r\n<li>sich mit einer Beschwerde an eine Aufsichtsbeh&ouml;rde zu wenden (Art. 15 Abs. 1 lit. f DSGVO). Eine Liste der Aufsichtsbeh&ouml;rden l&auml;sst sich unter&nbsp;<a href=\"https://www.bfdi.bund.de/DE/Infothek/Anschriften_Links/anschriften_links-node.html\">diesem Link</a>&nbsp;abrufen.</li>\r\n</ul>\r\n<h2>&Auml;nderung unserer Datenschutzerkl&auml;rung</h2>\r\n<p>Wir behalten uns vor, diese Datenschutzerkl&auml;rung ohne vorherige Ank&uuml;ndigung an neue rechtliche oder technische Sachverhalte sowie an &Auml;nderungen unserer Leistungen oder Prozesse anzupassen. Es gilt die jeweils auf der Onlinepr&auml;senz ver&ouml;ffentlichte Version der Datenschutzerkl&auml;rung.</p>\r\n</div>")
	if err := dao.SaveRecord(model_datenschutz); err != nil {
		fmt.Println(err)
	}

	model_b_person := models.NewRecord(collection)
	model_b_person.Set("Titel", "Begriffserklärung: Person")
	if err := dao.SaveRecord(model_b_person); err != nil {
		fmt.Println(err)
	}

	model_b_tt := models.NewRecord(collection)
	model_b_tt.Set("Titel", "Begriffserklärung: Titel-/Textanfänge")
	if err := dao.SaveRecord(model_b_tt); err != nil {
		fmt.Println(err)
	}

	model_kontakt := models.NewRecord(collection)
	model_kontakt.Set("Titel", "Kontakt")
	if err := dao.SaveRecord(model_kontakt); err != nil {
		fmt.Println(err)
	}

	model_recherche := models.NewRecord(collection)
	model_recherche.Set("Titel", "Recherche")
	if err := dao.SaveRecord(model_recherche); err != nil {
		fmt.Println(err)
	}
	return nil
}
